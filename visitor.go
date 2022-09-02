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
		// 栈顶开始查找
		for i := v.stack.Size() - 1; i > 0; i-- {
			f := v.stack.Get(i).(*stack.Frame)

			// 函数是一等公民时，需要根据变量的作用域进行判断
			funcObject, ok := frame.GetObject().(*stack.FuncObject)

			// 新写入的栈帧的 parent 与当前的 parent 相同时
			if f.GetScope().GetEncloseScope() == frame.GetScope().GetEncloseScope() {
				frame.SetParent(f.GetParent())
				break
			} else if f.GetScope() == frame.GetScope().GetEncloseScope() {
				// 新写入的栈帧是某个已有的栈帧的下级
				frame.SetParent(f)
				break
			} else if ok {
				referenceVariable := funcObject.GetReferenceVariable()
				if referenceVariable != nil && referenceVariable.GetEncloseScope() == f.GetScope() {
					frame.SetParent(f)
					break
				}
			}

		}

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

	// 闭包查询，闭包的变量不在父 scope 中，得一级一级的栈帧中查找。
	if object == nil {
		frame = v.stack.Peek().(*stack.Frame)
		for frame != nil {
			if frame.ContainsVariable(variable) {
				object = frame.GetObject()
				break
			}
			frame = frame.GetParent()
		}
	}
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
	case *parser.StmIfElseContext:
		return v.VisitStmIfElse(ctx)
	case *parser.StmReturnContext:
		return v.VisitStmReturn(ctx)
	case *parser.StmExprContext:
		return v.VisitStmExpr(ctx)
	case *parser.StmForContext:
		return v.VisitStmFor(ctx)
	case *parser.BlockFuncContext:
		return nil
	case *parser.ExprContext:
		return v.VisitExpr(ctx)
	case *parser.PrimaryContext:
		return v.VisitPrimary(ctx)
	case *parser.LiteralContext:
		return v.VisitLiteral(ctx)
	case *parser.BlockClassDeclarContext:
		return nil
	case *parser.ClassDeclarationContext:
		return v.VisitClassDeclaration(ctx)
	case *parser.ClassBodyContext:
		return v.VisitClassBody(ctx)
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
		leftObject := val1
		rightObject := val2
		switch val1.(type) {
		case *LeftValue:
			leftObject = val1.(*LeftValue).GetValue()
		}
		switch val2.(type) {
		case *LeftValue:
			rightObject = val2.(*LeftValue).GetValue()
		}

		//推导出来的该节点类型
		deriveType := v.at.GetTypeOfNode()[ctx]
		type1 := v.at.GetTypeOfNode()[ctx.Expr(0)]
		type2 := v.at.GetTypeOfNode()[ctx.Expr(1)]

		switch ctx.GetBop().GetTokenType() {
		case parser.GScriptParserMULT:
			if deriveType == sym.Int {
				return leftObject.(int) * rightObject.(int)
			} else if deriveType == sym.Float {
				return leftObject.(float64) * rightObject.(float64)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserDIV:
			if deriveType == sym.Int {
				return leftObject.(int) / rightObject.(int)
			} else if deriveType == sym.Float {
				return leftObject.(float64) / rightObject.(float64)
			} else {
				// todo crossoverJie 运行时错误
			}

		case parser.GScriptParserPLUS:
			if deriveType == sym.String {
				return fmt.Sprintf("%v", leftObject) + fmt.Sprintf("%v", rightObject)
			} else if deriveType == sym.Int {
				return leftObject.(int) + rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) + sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserSUB:
			if deriveType == sym.Int {
				return leftObject.(int) - rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) - sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}

		case parser.GScriptParserMOD:
			return leftObject.(int) % rightObject.(int)
		case parser.GScriptParserGT:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				// todo crossoverJie 运行时错误 字符串不能比较
			} else if deriveType == sym.Int {
				return leftObject.(int) > rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) > sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserLT:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				// todo crossoverJie 运行时错误 字符串不能比较
			} else if deriveType == sym.Int {
				return leftObject.(int) < rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) < sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserGE:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				// todo crossoverJie 运行时错误 字符串不能比较
			} else if deriveType == sym.Int {
				return leftObject.(int) >= rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) >= sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserLE:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				// todo crossoverJie 运行时错误 字符串不能比较
			} else if deriveType == sym.Int {
				return leftObject.(int) <= rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) <= sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserEQUAL:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				return fmt.Sprintf("%v", leftObject) == fmt.Sprintf("%v", rightObject)
			} else if deriveType == sym.Int {
				return leftObject.(int) == rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) == sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserNOTEQUAL:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				return fmt.Sprintf("%v", leftObject) != fmt.Sprintf("%v", rightObject)
			} else if deriveType == sym.Int {
				return leftObject.(int) != rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) != sym.Value2Float(rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserASSIGN:
			l := v.Visit(ctx.GetLhs()).(*LeftValue)
			r := val2
			switch val2.(type) {
			case *LeftValue:
				r = val2.(*LeftValue).GetValue()
			}
			// e = e+10
			l.SetValue(r)
			return r

		}

	}
	if ctx.GetBop() != nil && ctx.GetBop().GetTokenType() == parser.GScriptParserDOT {
		l := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext))
		switch l.(type) {
		case *LeftValue:
			left := l.(*LeftValue)
			switch left.GetValue().(type) {
			case *stack.ClassObject:
				classObject := left.GetValue().(*stack.ClassObject)
				if ctx.IDENTIFIER() != nil {
					_ = v.at.GetSymbolOfNode()[ctx.Expr(0)].(*sym.Variable)
					// todo crossoverJie this/super 从父级查找

					variable := v.at.GetSymbolOfNode()[ctx].(*sym.Variable)

					// person.age; 返回的是 age 的左值
					return NewLeftValue(variable, classObject)
				} else if ctx.FunctionCall() != nil {
					// person.getAge();
					// todo crossoverJie isSuper 赋值
					return v.receiveFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext), classObject, false)
				}
			}

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

// VisitFunctionCall 函数调用
func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	var ret interface{}
	name := ctx.IDENTIFIER().GetText()
	// todo crossoverJie 内置函数校验
	if name == "print" {
		v.print(ctx)
		return ret
	} else if name == "assertEqual" {
		v.assertEqual(ctx)
		return ret
	}

	// 默认构造函数
	symbol := v.at.GetSymbolOfNode()[ctx]
	switch symbol.(type) {
	case *sym.DefaultConstructorFunc:
		class := symbol.(*sym.DefaultConstructorFunc).GetClass()
		return v.initClassObject(class)
	}

	functionObject := v.getFunctionObject(ctx)

	// 如果对象的构造函数 Person(10)
	if functionObject.GetFunction().IsConstructor() {
		// 获取当前函数所归属的 class
		classObject := v.initClassObject(functionObject.GetFunction().GetEncloseScope().(*sym.Class))
		v.receiveFunctionCall(ctx, classObject, false)
		return classObject
	}

	// 构建函数调用的参数值
	paramValues := v.buildParamValues(ctx)

	// 执行函数调用
	ret = v.executeFunctionCall(functionObject, paramValues)

	// todo crossoverJie 支持 return
	return ret
}

// 初始化 classObject 对象
func (v *Visitor) initClassObject(class *sym.Class) *stack.ClassObject {
	object := stack.NewClassObject(class)

	var tempStack stack.Stack
	tempStack.Push(class)

	v.pushStack(stack.NewClassStackFrame(object))
	// todo crossoverJie 如果有父类需要一次初始化
	for !tempStack.IsEmpty() {
		pop := tempStack.Pop().(*sym.Class)
		v.initClassObjectField(pop, object)
	}
	v.popStack()
	return object
}

// 初始化 classObject 中的变量数据
func (v *Visitor) initClassObjectField(class *sym.Class, object *stack.ClassObject) {
	for _, symbol := range class.GetSymbols() {
		switch symbol.(type) {
		case *sym.Variable:
			// 为 class 中的变量初始化为空值
			object.SetValue(symbol.(*sym.Variable), nil)
		}
	}

	// 初始化变量，比如 class X{**int a=10**}
	ctx := class.GetCtx().(*parser.ClassDeclarationContext)
	v.VisitClassDeclaration(ctx)

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
	case *sym.Variable:
		// symbol 是函数变量类型
		variable := symbol.(*sym.Variable)
		value := v.getLeftValue(variable).GetValue()
		functionObject, ok := value.(*stack.FuncObject)
		if ok {
			function = functionObject.GetFunction()
			return functionObject
		}

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

// 类函数调用
func (v *Visitor) receiveFunctionCall(ctx *parser.FunctionCallContext, classObject *stack.ClassObject, isSuper bool) interface{} {

	// 拿到该 ctx 的functionObject
	v.pushStack(stack.NewClassStackFrame(classObject)) //这个入栈是为了类变量是函数
	funcObject := v.getFunctionObject(ctx)
	v.popStack()

	// todo crossoverJie 父类构造函数调用

	paramValues := v.buildParamValues(ctx)

	v.pushStack(stack.NewClassStackFrame(classObject))
	ret := v.executeFunctionCall(funcObject, paramValues)
	v.popStack()

	return ret
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	return v.VisitFunctionBody(ctx.FunctionBody().(*parser.FunctionBodyContext))
}

func (v *Visitor) VisitFunctionBody(ctx *parser.FunctionBodyContext) interface{} {
	return v.VisitBlock(ctx.Block().(*parser.BlockContext))
}

func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	if ctx.ClassBody() != nil {
		return v.VisitClassBody(ctx.ClassBody().(*parser.ClassBodyContext))
	}
	return nil
}

func (v *Visitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllClassBodyDeclaration() {
		ret = v.VisitClassBodyDeclaration(context.(*parser.ClassBodyDeclarationContext))
	}
	return ret
}

func (v *Visitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	return v.VisitMemberDeclaration(ctx.MemberDeclaration().(*parser.MemberDeclarationContext))
}

func (v *Visitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	if ctx.FieldDeclaration() != nil {
		return v.VisitFieldDeclaration(ctx.FieldDeclaration().(*parser.FieldDeclarationContext))
	}
	return nil
}

func (v *Visitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	return v.VisitVariableDeclarators(ctx.VariableDeclarators().(*parser.VariableDeclaratorsContext))
}

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
		case *sym.Func:
			// 函数变量
			funcObject := stack.NewFuncObject(symbol.(*sym.Func))
			ret = funcObject
		}
		return ret
	}
	return nil
}

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
		ret := v.Visit(ctx.Expr())
		switch ret.(type) {
		case *LeftValue:
			ret = ret.(*LeftValue).GetValue()
		case *stack.FuncObject:
			funcObject := ret.(*stack.FuncObject)
			v.setClosureValues(funcObject)
		case *stack.ClassObject:
			//classObject := ret.(*stack.ClassObject)
		}
		// todo crossoverJie class 类的闭包

		// todo crossoverJie 支持 return
		return ret
	}
	return nil
}

// 将闭包变量复制到当前 functionObject 中，同时赋值。
func (v *Visitor) setClosureValues(funcObject *stack.FuncObject) {
	if funcObject.GetFunction().GetEncloseScope() == nil {
		return
	}

	closureVar := funcObject.GetFunction().GetClosureVar()
	for _, val := range closureVar.List() {
		variable := val.(*sym.Variable)
		leftValue := v.getLeftValue(variable)
		funcObject.SetValue(variable, leftValue.GetValue())
	}
}

// 将闭包变量复制到当前的 classObject 中
func (v *Visitor) setClassClosureValues(classObject *stack.ClassObject) {
	//object := stack.NewEmptyObject()
	//for variable, v := range classObject.AllField() {
	//	if v!=nil {
	//
	//	}
	//}
}

func (v *Visitor) VisitStmExpr(ctx *parser.StmExprContext) interface{} {
	return v.Visit(ctx.GetStatementExpression())
}

func (v *Visitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return v.Visit(ctx.Expr())
}
