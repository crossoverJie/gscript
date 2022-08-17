package gscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
	"github.com/crossoverJie/gscript/stack"
	sym "github.com/crossoverJie/gscript/symbol"
	"strconv"
)

type Visitor struct {
	parser.BaseGScriptVisitor
	at    *resolver.AnnotatedTree
	stack stack.Stack
}

func NewVisitor(at *resolver.AnnotatedTree) *Visitor {
	return &Visitor{at: at}
}

func ArithmeticOperators(script string) interface{} {
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	tree := parser.NewGScriptParser(stream).Prog()
	visitor := Visitor{}
	return visitor.Visit(tree)
}

/**
栈帧入栈
*/
func (v *Visitor) pushStack(frame *stack.Frame) {
	// todo crossoverJie parentFrame 设置
	if !v.stack.IsEmpty() {
		if frame.GetParent() == nil {
			// 上一级作用域作为本级作用域的父级
			frame.SetParent(v.stack.Peek().(*stack.Frame))
		}

	}

	v.stack.Push(frame)
}

func (v *Visitor) popStack() {
	v.stack.Pop()
}

// 在整个栈帧中获取左值
func (v *Visitor) getLeftValue(variable *sym.Variable) *LeftValue {
	frame := v.stack.Peek().(*stack.Frame)
	var object stack.Object
	for frame != nil {
		// 按照作用域获取变量值，内部作用域覆盖外部作用域。
		if frame.GetScope().ContainsSymbol(variable) {
			object = frame.GetObject()
			break
		}
		frame = frame.GetParent()
	}

	// todo crossoverJie 闭包查询
	return NewLeftValue(variable, object)
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(ctx)
	case *parser.BlockStmContext:
		return v.VisitBlockStm(ctx)
	case *parser.StmBlockLabelContext:
		return v.VisitStmBlockLabel(ctx)
	case *parser.BlockVarDeclarContext:
		return v.VisitBlockVarDeclar(ctx)
	case *parser.VariableDeclaratorsContext:
		return v.VisitVariableDeclarators(ctx)
	case *parser.ParseContext:
		return v.VisitParse(ctx)
	//case *parser.PostfixExprContext:
	//	return v.VisitPostfixExpr(ctx)
	//case *parser.NotExprContext:
	//	return v.VisitNotExpr(ctx)
	//case *parser.MultDivExprContext:
	//	return v.VisitMultDivExpr(ctx)
	//case *parser.PrimaryExprContext:
	//	return v.VisitPrimaryExpr(ctx)
	//case *parser.LiterPrimaryContext:
	//	return v.VisitLiterPrimary(ctx)
	//case *parser.IdentifierPrimaryContext:
	//	return v.VisitIdentifierPrimary(ctx)
	//case *parser.IntContext:
	//	return v.VisitInt(ctx)
	//case *parser.FloatContext:
	//	return v.VisitFloat(ctx)
	//case *parser.StringContext:
	//	return v.VisitString(ctx)
	//case *parser.BoolContext:
	//	return v.VisitBool(ctx)
	//case *parser.NullContext:
	//	return v.VisitNull(ctx)
	//case *parser.PlusSubExprContext:
	//	return v.VisitPlusSubExpr(ctx)
	//case *parser.ExprPrimaryContext:
	//	return v.VisitExprPrimary(ctx)
	//case *parser.ModExprContext:
	//	return v.VisitModExpr(ctx)
	//case *parser.GLeExprContext:
	//	return v.VisitGLeExpr(ctx)
	//case *parser.EqualOrNotContext:
	//	return v.VisitEqualOrNot(ctx)
	case *parser.StmIfElseContext:
		return v.VisitStmIfElse(ctx)
	case *parser.StmReturnContext:
		return v.VisitStmReturn(ctx)
	case *parser.StmExprContext:
		return v.VisitStmExpr(ctx)
	case *parser.StmForContext:
		return v.VisitStmFor(ctx)
	//case *parser.AssignExprContext:
	//	return v.VisitAssignExpr(ctx)
	//case *parser.FuncCallExprContext:
	//	return v.VisitFuncCallExpr(ctx)
	case *parser.BlockFuncContext:
		return v.VisitBlockFunc(ctx)
	case *parser.ExprContext:
		return v.VisitExpr(ctx)
	case *parser.PrimaryContext:
		return v.VisitPrimary(ctx)
	case *parser.LiteralContext:
		return v.VisitLiteral(ctx)
	default:
		panic("Unknown context")
	}
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) interface{} {
	// 将scope写入栈帧
	scope := v.at.GetNode2Scope()[ctx]
	v.pushStack(stack.NewBlockScopeFrame(scope))
	ret := v.VisitBlockStms(ctx.BlockStatements().(*parser.BlockStmsContext))
	v.popStack()
	return ret
}

func (v *Visitor) VisitBlockStms(ctx *parser.BlockStmsContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllBlockStatement() {
		ret = v.Visit(context)
	}
	return ret
}

func (v *Visitor) VisitBlockVarDeclar(ctx *parser.BlockVarDeclarContext) interface{} {
	return v.Visit(ctx.VariableDeclarators())
}

func (v *Visitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllVariableDeclarator() {
		ret = v.VisitVariableDeclarator(context.(*parser.VariableDeclaratorContext))
	}
	return ret
}

func (v *Visitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	var ret interface{}
	leftValue := v.VisitVariableDeclaratorId(ctx.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext)).(*LeftValue)
	if ctx.VariableInitializer() != nil {
		ret = v.VisitVariableInitializer(ctx.VariableInitializer().(*parser.VariableInitializerContext))
		switch ret.(type) {
		case *LeftValue:
			ret = ret.(*LeftValue).GetValue()
		}
		// 为变量赋值
		// int e=10   int e = foo()
		leftValue.SetValue(ret)
	}
	return ret
}

func (v *Visitor) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	symbol := v.at.GetSymbolOfNode()[ctx]
	value := v.getLeftValue(symbol.(*sym.Variable))
	return value
}
func (v *Visitor) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	// todo array init
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	return nil
}

func (v *Visitor) VisitBlockStm(ctx *parser.BlockStmContext) interface{} {
	return v.Visit(ctx.Statement())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// 将 scope 写入栈帧
	scope := v.at.GetNode2Scope()[ctx]
	if scope != nil {
		v.pushStack(stack.NewBlockScopeFrame(scope))
	}

	ret := v.VisitBlockStms(ctx.BlockStatements().(*parser.BlockStmsContext))

	if scope != nil {
		v.popStack()
	}
	return ret
}

func (v *Visitor) VisitStmBlockLabel(ctx *parser.StmBlockLabelContext) interface{} {
	return v.VisitBlock(ctx.GetBlockLabel().(*parser.BlockContext))
}

func (v *Visitor) VisitParse(ctx *parser.ParseContext) interface{} {
	for _, expr := range ctx.GetExpr_list() {
		return v.Visit(expr)
	}
	return nil
}

func (v *Visitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	var ret interface{}
	if ctx.Primary() != nil {
		ret = v.Visit(ctx.Primary())
	}

	if ctx.GetBop() != nil && len(ctx.AllExpr()) >= 2 {
		val1 := v.Visit(ctx.GetLhs())
		val2 := v.Visit(ctx.GetRhs())
		var (
			v1i int
			v1f float64
			v2i int
			v2f float64
		)
		switch val1.(type) {
		case int:
			v1i = val1.(int)
		case float64:
			v1f = val1.(float64)
		case *LeftValue:
			v1i = val1.(*LeftValue).GetValue().(int)
		}

		switch val2.(type) {
		case int:
			v2i = val2.(int)
		case float64:
			v2f = val2.(float64)
		case *LeftValue:
			v2i = val2.(*LeftValue).GetValue().(int)
		}
		switch ctx.GetBop().GetTokenType() {
		case parser.GScriptParserMULT:
			if v1i != 0 && v2i != 0 {
				return v1i * v2i
			}
			if v1i != 0 && v2f != 0 {
				return float64(v1i) * v2f
			}
			if v1f != 0 && v2i != 0 {
				return v1f * float64(v2i)
			}
			if v1f != 0 && v2f != 0 {
				return v1f * v2f
			}
		case parser.GScriptParserDIV:
			if v1i != 0 && v2i != 0 {
				return v1i / v2i
			}
			if v1i != 0 && v2f != 0 {
				return float64(v1i) / v2f
			}
			if v1f != 0 && v2i != 0 {
				return v1f / float64(v2i)
			}
			if v1f != 0 && v2f != 0 {
				return v1f / v2f
			}

		case parser.GScriptParserPLUS:
			if v1i != 0 && v2i != 0 {
				return v1i + v2i
			}
			if v1i != 0 && v2f != 0 {
				return float64(v1i) + v2f
			}
			if v1f != 0 && v2i != 0 {
				return v1f + float64(v2i)
			}
			if v1f != 0 && v2f != 0 {
				return v1f + v2f
			}
		case parser.GScriptParserSUB:
			if v1i != 0 && v2i != 0 {
				return v1i - v2i
			}
			if v1i != 0 && v2f != 0 {
				return float64(v1i) - v2f
			}
			if v1f != 0 && v2i != 0 {
				return v1f - float64(v2i)
			}
			if v1f != 0 && v2f != 0 {
				return v1f - v2f
			}

		case parser.GScriptParserMOD:
			lhs := v.Visit(ctx.GetLhs())
			rhs := v.Visit(ctx.GetRhs())
			return lhs.(int) % rhs.(int)
		case parser.GScriptParserGT:
			return v1i > v2i
		case parser.GScriptParserLT:
			return v1i < v2i
		case parser.GScriptParserGE:
			return v1i >= v2i
		case parser.GScriptParserLE:
			return v1i <= v2i
		case parser.GScriptParserEQUAL:
			return v1i == v2i
		case parser.GScriptParserNOTEQUAL:
			return v1i != v2i
		case parser.GScriptParserASSIGN:
			l := v.Visit(ctx.GetLhs()).(*LeftValue)
			r := val2
			switch val2.(type) {
			case *LeftValue:
				r = val2.(*LeftValue).GetValue()
			}
			// e = e+10
			l.SetValue(r)
		}

	}
	if ctx.GetPostfix() != nil {
		lhs := ctx.GetLhs()
		value := v.Visit(lhs)
		switch value.(type) {
		case *LeftValue:
			leftValue := value.(*LeftValue)
			switch ctx.GetPostfix().GetTokenType() {
			case parser.GScriptParserINC:
				leftValue.SetValue(leftValue.GetValue().(int) + 1)
				return value
			case parser.GScriptParserDEC:
				leftValue.SetValue(leftValue.GetValue().(int) - 1)
				return value
			}
		case int:
			switch ctx.GetPostfix().GetTokenType() {
			case parser.GScriptParserINC:
				value = value.(int) + 1
				return value

			case parser.GScriptParserDEC:
				value = value.(int) - 1
				return value
			}
		}
	}

	if ctx.GetPrefix() != nil {
		rhs := ctx.GetRhs()
		value := v.Visit(rhs)
		switch value.(type) {
		case bool:
			return !value.(bool)
		}
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		panic(fmt.Sprintf("invalid ! symbol in line:%d and column:%d", line, column))
	}

	if ctx.FunctionCall() != nil {
		return v.VisitFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
	}

	return ret
}

//func (v *Visitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
//	return v.Visit(ctx.Primary())
//}

//func (v *Visitor) VisitPostfixExpr(ctx *parser.PostfixExprContext) interface{} {
//	lhs := ctx.GetLhs()
//	value := v.Visit(lhs)
//	switch value.(type) {
//	case *LeftValue:
//		leftValue := value.(*LeftValue)
//		switch ctx.GetPostfix().GetTokenType() {
//		case parser.GScriptParserINC:
//			leftValue.SetValue(leftValue.GetValue().(int) + 1)
//			return value
//		case parser.GScriptParserDEC:
//			leftValue.SetValue(leftValue.GetValue().(int) - 1)
//			return value
//		}
//	case int:
//		switch ctx.GetPostfix().GetTokenType() {
//		case parser.GScriptParserINC:
//			value = value.(int) + 1
//			return value
//
//		case parser.GScriptParserDEC:
//			value = value.(int) - 1
//			return value
//		}
//	}
//
//	panic("invalid postfix")
//}

//func (v *Visitor) VisitFuncCallExpr(ctx *parser.FuncCallExprContext) interface{} {
//	return v.VisitFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
//}

// VisitFunctionCall 函数调用
func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	var ret interface{}
	//name := ctx.IDENTIFIER().GetText()
	// todo crossoverJie 内置函数校验

	//symbol := v.at.GetSymbolOfNode()[ctx]
	// todo crossoverJie 默认构造函数

	functionObject := v.getFunctionObject(ctx)

	// todo crossoverJie 如果对象的构造函数

	// 构建函数调用的参数值
	paramValues := v.buildParamValues(ctx)

	// 执行函数调用
	ret = v.executeFunctionCall(functionObject, paramValues)
	return ret
}

// 获取函数的 object 对象，需要用来压栈
func (v *Visitor) getFunctionObject(ctx *parser.FunctionCallContext) *stack.FuncObject {
	var (
		funcObject *stack.FuncObject
		function   *sym.Func
	)
	symbol := v.at.GetSymbolOfNode()[ctx]
	switch symbol.(type) {
	case *sym.Func:
		function = symbol.(*sym.Func)
		// todo crossoverJie symbol 是函数变量类型
	default:
		// todo crossoverJie 函数提示
		name := ctx.IDENTIFIER().GetText()
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		panic(fmt.Sprintf("unable find func %s in line:%d and column:%d", name, line, column))

	}

	funcObject = stack.NewFuncObject(function)
	return funcObject
}

// 构建函数调用的参数值 myfunc(2+2+a) 2+2+a 的值
func (v *Visitor) buildParamValues(ctx *parser.FunctionCallContext) []interface{} {
	ret := make([]interface{}, 0)
	if ctx.ExpressionList() == nil {
		return ret
	}
	for _, context := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpr() {
		value := v.Visit(context)
		switch value.(type) {
		case *LeftValue:
			ret = append(ret, value.(*LeftValue).GetValue())
		default:
			ret = append(ret, value)
		}
	}
	return ret
}

// 执行函数调用
// 1. 先给参数赋值
// 2. 调用函数，内部就是递归调用 block
func (v *Visitor) executeFunctionCall(functionObject *stack.FuncObject, paramValues []interface{}) interface{} {
	var ret interface{}

	// 添加函数栈帧
	frame := stack.NewObjectStackFrame(functionObject)
	v.pushStack(frame)

	// 拿到函数声明时所对应的 ctx, type_scope_resolver.go EnterFunctionDeclaration, g4:functionDeclaration
	declarationContext := functionObject.GetFunction().GetCtx().(*parser.FunctionDeclarationContext)
	formalParametersContext := declarationContext.FormalParameters().(*parser.FormalParametersContext)
	if formalParametersContext.FormalParameterList() != nil {
		// 为函数参数赋值
		for i, context := range formalParametersContext.FormalParameterList().(*parser.FormalParameterListContext).AllFormalParameter() {
			parameterContext := context.(*parser.FormalParameterContext)
			value := v.VisitVariableDeclaratorId(parameterContext.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext))
			switch value.(type) {
			case *LeftValue:
				leftValue := value.(*LeftValue)
				leftValue.SetValue(paramValues[i])
			}
		}
	}

	// 执行方法体
	ret = v.VisitFunctionDeclaration(declarationContext)

	v.popStack()
	return ret
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	return v.VisitFunctionBody(ctx.FunctionBody().(*parser.FunctionBodyContext))
}

func (v *Visitor) VisitFunctionBody(ctx *parser.FunctionBodyContext) interface{} {
	return v.VisitBlock(ctx.Block().(*parser.BlockContext))
}

//func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
//	rhs := ctx.GetRhs()
//	value := v.Visit(rhs)
//	switch value.(type) {
//	case bool:
//		return !value.(bool)
//	}
//	line := ctx.GetStart().GetLine()
//	column := ctx.GetStart().GetColumn()
//	panic(fmt.Sprintf("invalid ! symbol in line:%d and column:%d", line, column))
//}

//func (v *Visitor) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
//	val1 := v.Visit(ctx.GetLhs())
//	val2 := v.Visit(ctx.GetRhs())
//
//	var (
//		v1i int
//		v1f float64
//		v2i int
//		v2f float64
//	)
//	switch val1.(type) {
//	case int:
//		v1i = val1.(int)
//	case float64:
//		v1f = val1.(float64)
//	}
//
//	switch val2.(type) {
//	case int:
//		v2i = val2.(int)
//	case float64:
//		v2f = val2.(float64)
//	}
//
//	if ctx.GetBop().GetTokenType() == parser.GScriptParserMULT {
//		if v1i != 0 && v2i != 0 {
//			return v1i * v2i
//		}
//		if v1i != 0 && v2f != 0 {
//			return float64(v1i) * v2f
//		}
//		if v1f != 0 && v2i != 0 {
//			return v1f * float64(v2i)
//		}
//		if v1f != 0 && v2f != 0 {
//			return v1f * v2f
//		}
//	} else {
//		if v1i != 0 && v2i != 0 {
//			return v1i / v2i
//		}
//		if v1i != 0 && v2f != 0 {
//			return float64(v1i) / v2f
//		}
//		if v1f != 0 && v2i != 0 {
//			return v1f / float64(v2i)
//		}
//		if v1f != 0 && v2f != 0 {
//			return v1f / v2f
//		}
//	}
//	return 0
//}

//func (v *Visitor) VisitPlusSubExpr(ctx *parser.PlusSubExprContext) interface{} {
//	val1 := v.Visit(ctx.GetLhs())
//	val2 := v.Visit(ctx.GetRhs())
//
//	var (
//		v1i int
//		v1f float64
//		v2i int
//		v2f float64
//	)
//	switch val1.(type) {
//	case int:
//		v1i = val1.(int)
//	case float64:
//		v1f = val1.(float64)
//	case *LeftValue:
//		value := val1.(*LeftValue).GetValue()
//		switch value.(type) {
//		case int:
//			v1i = value.(int)
//		case float64:
//			v1f = value.(float64)
//		}
//	}
//
//	switch val2.(type) {
//	case int:
//		v2i = val2.(int)
//	case float64:
//		v2f = val2.(float64)
//	}
//
//	// todo crossoverJie v1i 值为0的情况
//	if ctx.GetBop().GetTokenType() == parser.GScriptLexerPLUS {
//		if v1i != 0 && v2i != 0 {
//			return v1i + v2i
//		}
//		if v1i != 0 && v2f != 0 {
//			return float64(v1i) + v2f
//		}
//		if v1f != 0 && v2i != 0 {
//			return v1f + float64(v2i)
//		}
//		if v1f != 0 && v2f != 0 {
//			return v1f + v2f
//		}
//	} else {
//		if v1i != 0 && v2i != 0 {
//			return v1i - v2i
//		}
//		if v1i != 0 && v2f != 0 {
//			return float64(v1i) - v2f
//		}
//		if v1f != 0 && v2i != 0 {
//			return v1f - float64(v2i)
//		}
//		if v1f != 0 && v2f != 0 {
//			return v1f - v2f
//		}
//	}
//	return 0
//}

//func (v *Visitor) VisitLiterPrimary(ctx *parser.LiterPrimaryContext) interface{} {
//	return v.Visit(ctx.Literal())
//}

//func (v *Visitor) VisitIdentifierPrimary(ctx *parser.IdentifierPrimaryContext) interface{} {
//	var ret interface{}
//	//fmt.Println(ctx.IDENTIFIER().GetText())
//	//if ctx.IDENTIFIER() != nil {
//	symbol := v.at.GetSymbolOfNode()[ctx]
//	switch symbol.(type) {
//	case *sym.Variable:
//		ret = v.getLeftValue(symbol.(*sym.Variable))
//	}
//	//}
//	return ret
//}

//func (v *Visitor) VisitInt(ctx *parser.IntContext) interface{} {
//	val, _ := strconv.Atoi(ctx.GetText())
//	return val
//}
//
//func (v *Visitor) VisitFloat(ctx *parser.FloatContext) interface{} {
//	val, _ := strconv.ParseFloat(ctx.GetText(), 0)
//	return val
//}
//
//func (v *Visitor) VisitString(ctx *parser.StringContext) interface{} {
//	return ctx.GetText()[1 : len(ctx.GetText())-1]
//}
//
//func (v *Visitor) VisitBool(ctx *parser.BoolContext) interface{} {
//	val, _ := strconv.ParseBool(ctx.GetText())
//	return val
//}
//
//func (v *Visitor) VisitNull(ctx *parser.NullContext) interface{} {
//	return nil
//}

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.DECIMAL_LITERAL() != nil {
		val, _ := strconv.Atoi(ctx.GetText())
		return val
	}
	if ctx.FLOAT_LITERAL() != nil {
		val, _ := strconv.ParseFloat(ctx.GetText(), 0)
		return val
	}
	if ctx.STRING_LITERAL() != nil {
		return ctx.GetText()[1 : len(ctx.GetText())-1]
	}
	if ctx.BOOL_LITERAL() != nil {
		val, _ := strconv.ParseBool(ctx.GetText())
		return val
	}
	if ctx.NULL_LITERAL() != nil {
		return nil
	}
	return nil
}

//func (v *Visitor) VisitExprPrimary(ctx *parser.ExprPrimaryContext) interface{} {
//	return v.Visit(ctx.Expr())
//}

func (v *Visitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	// 字面量
	if ctx.Literal() != nil {
		return v.Visit(ctx.Literal())
	}

	// 标识符
	if ctx.IDENTIFIER() != nil {
		var ret interface{}
		//fmt.Println(ctx.IDENTIFIER().GetText())
		symbol := v.at.GetSymbolOfNode()[ctx]
		switch symbol.(type) {
		case *sym.Variable:
			ret = v.getLeftValue(symbol.(*sym.Variable))
		}
		return ret
	}
	return nil
}

//
//func (v *Visitor) VisitModExpr(ctx *parser.ModExprContext) interface{} {
//	lhs := v.Visit(ctx.GetLhs())
//	rhs := v.Visit(ctx.GetRhs())
//	return lhs.(int) % rhs.(int)
//}

//func (v *Visitor) VisitGLeExpr(ctx *parser.GLeExprContext) interface{} {
//	var left, right int
//	l := v.Visit(ctx.Expr(0))
//	switch l.(type) {
//	case *LeftValue:
//		left = l.(*LeftValue).GetValue().(int)
//	case int:
//		left = l.(int)
//	}
//	r := v.Visit(ctx.Expr(1))
//	switch r.(type) {
//	case *LeftValue:
//		right = r.(*LeftValue).GetValue().(int)
//	case int:
//		right = r.(int)
//	}
//
//	switch ctx.GetBop().GetTokenType() {
//	case parser.GScriptParserGT:
//		return left > right
//	case parser.GScriptParserLT:
//		return left < right
//	case parser.GScriptParserGE:
//		return left >= right
//	case parser.GScriptParserLE:
//		return left <= right
//	default:
//		panic("invalid symbol")
//	}
//}

//func (v *Visitor) VisitEqualOrNot(ctx *parser.EqualOrNotContext) interface{} {
//	left := v.Visit(ctx.Expr(0)).(int)
//	right := v.Visit(ctx.Expr(1)).(int)
//	switch ctx.GetBop().GetTokenType() {
//	case parser.GScriptParserEQUAL:
//		return left == right
//	case parser.GScriptParserNOTEQUAL:
//		return left != right
//	default:
//		panic("invalid symbol")
//	}
//}

//func (v *Visitor) VisitAssignExpr(ctx *parser.AssignExprContext) interface{} {
//	lhs := ctx.GetLhs()
//	rhs := ctx.GetRhs()
//	l := v.Visit(lhs).(*LeftValue)
//	r := v.Visit(rhs)
//	switch v.Visit(rhs).(type) {
//	case *LeftValue:
//		r = v.Visit(rhs).(*LeftValue).GetValue()
//	}
//	switch ctx.GetBop().GetTokenType() {
//	case parser.GScriptParserASSIGN:
//		l.SetValue(r)
//	}
//	return r
//
//}

func (v *Visitor) VisitStmIfElse(ctx *parser.StmIfElseContext) interface{} {
	condition := v.VisitParExpression(ctx.ParExpression().(*parser.ParExpressionContext)).(bool)
	if condition {
		return v.Visit(ctx.Statement(0))
	} else if ctx.ELSE() != nil {
		return v.Visit(ctx.Statement(1))
	}
	return nil
}

func (v *Visitor) VisitStmFor(ctx *parser.StmForContext) interface{} {
	scope := v.at.GetNode2Scope()[ctx]
	v.pushStack(stack.NewBlockScopeFrame(scope))

	forControl := ctx.ForControl()
	v.VisitForControl(forControl.(*parser.ForControlContext))

	v.popStack()
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForControl(ctx *parser.ForControlContext) interface{} {
	var ret interface{}
	forInit := ctx.ForInit()
	if forInit != nil {
		ret = v.VisitForInit(forInit.(*parser.ForInitContext))
	}
	for {
		condition := true
		if ctx.Expr() != nil {
			value := v.Visit(ctx.Expr())
			switch value.(type) {
			case *LeftValue:
				condition = value.(*LeftValue).GetValue().(bool)
			case bool:
				condition = value.(bool)
			}

			if !condition {
				break
			}
			ret = v.Visit(ctx.GetParent().(*parser.StmForContext).Statement())

			// todo crossoverJie break/return

			// i++
			if ctx.GetForUpdate() != nil {
				v.VisitExpressionList(ctx.GetForUpdate().(*parser.ExpressionListContext))
			}
		}
	}

	return ret
}

func (v *Visitor) VisitForInit(ctx *parser.ForInitContext) interface{} {
	var ret interface{}
	if ctx.VariableDeclarators() != nil {
		ret = v.VisitVariableDeclarators(ctx.VariableDeclarators().(*parser.VariableDeclaratorsContext))
	} else if ctx.ExpressionList() != nil {
		ret = v.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext))
	}

	return ret
}

func (v *Visitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllExpr() {
		ret = v.Visit(context)
	}
	return ret
}

func (v *Visitor) VisitStmReturn(ctx *parser.StmReturnContext) interface{} {
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	return nil
}

func (v *Visitor) VisitStmExpr(ctx *parser.StmExprContext) interface{} {
	return v.Visit(ctx.GetStatementExpression())
}

func (v *Visitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return v.Visit(ctx.Expr())
}
