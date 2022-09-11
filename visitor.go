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
	case *parser.StmBreakContext:
		return v.VisitStmBreak(ctx)
	case *parser.StmContinueContext:
		return v.VisitStmContinue(ctx)
	case *parser.StmExprContext:
		return v.VisitStmExpr(ctx)
	case *parser.StmForContext:
		return v.VisitStmFor(ctx)
	case *parser.StmWhileContext:
		return v.VisitStmWhile(ctx)
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
		switch ret.(type) {
		case *ContinueObject:
			return ret
		case *BreakObject:
			return ret
		case *ReturnObject:
			return ret
		}
		//ret = retTemp
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
		case *ArrayObject:
			// 数组赋值 int b=a[0];
			arrayObject := ret.(*ArrayObject)
			ret = arrayObject.GetIndexValue()
		}
		// 为变量赋值
		// int e=10   int e = foo()
		switch ret.(type) {
		case int:
			if leftValue.GetVariable().GetType() != sym.Int {
				// string a=10; 校验这类错误
				// todo crossoverJie 运行时错误，还有其他类型的校验，any 类型不需要校验
			}
		}
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
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	// array init
	if ctx.ArrayInitializer() != nil {
		allContext := ctx.ArrayInitializer().(*parser.ArrayInitializerContext)
		length := v.VisitArrayInitializer(ctx.ArrayInitializer().(*parser.ArrayInitializerContext))
		//var array []interface{}
		array := make([]interface{}, length.(int))
		for _, context := range allContext.AllVariableInitializer() {
			val := v.VisitVariableInitializer(context.(*parser.VariableInitializerContext))
			array = append(array, val)
		}
		return array
	}
	return nil
}

func (v *Visitor) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	if ctx.LBRACK() != nil && ctx.RBRACK() != nil {
		val, _ := strconv.Atoi(ctx.DECIMAL_LITERAL().GetText())
		return val
	}
	return 0
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

	// 获取数组数据
	if ctx.GetArray() != nil && ctx.GetIndex() != nil {
		left := v.VisitExpr(ctx.GetArray().(*parser.ExprContext))
		index := v.VisitExpr(ctx.GetIndex().(*parser.ExprContext))
		switch index.(type) {
		case int:
			ret = NewArrayObject(left.(*LeftValue), index.(int))
		case *LeftValue:
			leftValue := index.(*LeftValue)
			ret = NewArrayObject(left.(*LeftValue), leftValue.GetValue().(int))
		}

	}

	if ctx.GetBop() != nil && len(ctx.AllExpr()) >= 2 {
		val1 := v.Visit(ctx.GetLhs())
		val2 := v.Visit(ctx.GetRhs())
		leftObject := val1
		rightObject := val2
		switch val1.(type) {
		case *LeftValue:
			leftObject = val1.(*LeftValue).GetValue()
		case *ArrayObject:
			leftObject = val1.(*ArrayObject).GetIndexValue()
		}
		switch val2.(type) {
		case *LeftValue:
			rightObject = val2.(*LeftValue).GetValue()
		case *ArrayObject:
			rightObject = val2.(*ArrayObject).GetIndexValue()

		}

		//推导出来的该节点类型
		deriveType := v.at.GetTypeOfNode()[ctx]
		type1 := v.at.GetTypeOfNode()[ctx.Expr(0)]
		type2 := v.at.GetTypeOfNode()[ctx.Expr(1)]
		if deriveType == sym.Any {
			// 两个值都是any类型，需要运行时通过值判断
			deriveType = sym.GetUpperTypeWithValue(leftObject, rightObject)
		}

		switch ctx.GetBop().GetTokenType() {
		case parser.GScriptParserMULT:
			if deriveType == sym.Int {
				return leftObject.(int) * rightObject.(int)
			} else if deriveType == sym.Float {
				return leftObject.(float64) * rightObject.(float64)
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(type1, ctx.GetBop().GetTokenType(), leftObject, rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserDIV:
			if deriveType == sym.Int {
				return leftObject.(int) / rightObject.(int)
			} else if deriveType == sym.Float {
				return leftObject.(float64) / rightObject.(float64)
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(type1, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(type1, ctx.GetBop().GetTokenType(), leftObject, rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserSUB:
			if deriveType == sym.Int {
				return leftObject.(int) - rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) - sym.Value2Float(rightObject)
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(type1, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
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
			} else if deriveType == sym.Nil {
				if leftObject == nil && rightObject == nil {
					return true
				} else {
					return false
				}
			} else if deriveType == sym.Any {
				// 两个 any 值进行比较
				return leftObject == rightObject
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
			} else {
				return leftObject == rightObject
			}
		case parser.GScriptParserNOTEQUAL:
			deriveType = sym.GetUpperType(type1, type2)
			if deriveType == sym.String {
				return fmt.Sprintf("%v", leftObject) != fmt.Sprintf("%v", rightObject)
			} else if deriveType == sym.Int {
				return leftObject.(int) != rightObject.(int)
			} else if deriveType == sym.Float {
				return sym.Value2Float(leftObject) != sym.Value2Float(rightObject)
			} else if deriveType == sym.Nil {
				if leftObject != nil || rightObject != nil {
					return true
				} else {
					return false
				}
			} else if deriveType == sym.Any {
				// 两个 any 值进行比较
				return leftObject != rightObject
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				return v.callOpFunction(sym.Bool, ctx.GetBop().GetTokenType(), leftObject, rightObject)
			} else {
				// todo crossoverJie 运行时错误
			}
		case parser.GScriptParserASSIGN:
			switch val1.(type) {
			case *LeftValue:
				l := val1.(*LeftValue)
				r := val2
				switch val2.(type) {
				case *LeftValue:
					r = val2.(*LeftValue).GetValue()
				}
				// e = e+10
				l.SetValue(r)
				return r
			case *ArrayObject:
				// 数组赋值 a[1]=3;
				arrayObject := val1.(*ArrayObject)
				arrayObject.SetIndexValue(val2)
				//array,ok := arrayObject.GetLeftValue().GetValue().([]interface{})
				//if ok {
				//	array[arrayObject.GetIndex()] = val2
				//}
			}

		}

	}
	if ctx.GetBop() != nil && (ctx.GetBop().GetTokenType() == parser.GScriptParserAND || ctx.GetBop().GetTokenType() == parser.GScriptParserOR) {
		// &&
		left := v.VisitExpr(ctx.GetLhs().(*parser.ExprContext))
		right := v.VisitExpr(ctx.GetRhs().(*parser.ExprContext))
		var (
			leftCondition  = false
			rightCondition = false
		)

		switch left.(type) {
		case bool:
			leftCondition = left.(bool)
		case *LeftValue:
			b, ok := left.(*LeftValue).GetValue().(bool)
			if ok {
				leftCondition = b
			}
		}

		switch right.(type) {
		case bool:
			rightCondition = right.(bool)
		case *LeftValue:
			b, ok := right.(*LeftValue).GetValue().(bool)
			if ok {
				rightCondition = b
			}
		}
		if ctx.GetBop().GetTokenType() == parser.GScriptParserAND {
			return leftCondition && rightCondition
		} else if ctx.GetBop().GetTokenType() == parser.GScriptParserOR {
			return leftCondition || rightCondition
		} else {
			return false
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
			case *LeftValue:
				leftValue := left.GetValue().(*LeftValue)
				if ctx.IDENTIFIER() != nil {
					v1 := v.at.GetSymbolOfNode()[ctx].(*sym.Variable)
					switch leftValue.GetValue().(type) {
					case *stack.ClassObject:
						classObject := leftValue.GetValue().(*stack.ClassObject)
						return NewLeftValue(v1, classObject)
					}
				}
			case stack.Object:
				object := left.GetValue().(stack.Object)
				v1 := v.at.GetSymbolOfNode()[ctx].(*sym.Variable)
				value := object.GetValue(v1)
				return value
			}

		}
	}

	// 后缀计算
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

	// 前缀计算
	if ctx.GetPrefix() != nil {
		rhs := ctx.GetRhs()
		value := v.Visit(rhs)
		if ctx.GetPrefix().GetTokenType() == parser.GScriptParserBANG {
			switch value.(type) {
			case bool:
				return !value.(bool)
			}
			line := ctx.GetStart().GetLine()
			column := ctx.GetStart().GetColumn()
			panic(fmt.Sprintf("invalid ! symbol in line:%d and column:%d", line, column))
		} else if ctx.GetPrefix().GetTokenType() == parser.GScriptParserSUB {
			// int a = -10; int b=-10.1;
			switch value.(type) {
			case *LeftValue:
				getValue := value.(*LeftValue).GetValue()
				switch getValue.(type) {
				case int:
					return -getValue.(int)
				case float64:
					return -getValue.(float64)
				}
			case int:
				return -value.(int)
			case float64:
				return -value.(float64)
			}
		}

	}

	if ctx.FunctionCall() != nil {
		return v.VisitFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
	}

	return ret
}

// 执行自定义的运算符重载函数
func (v *Visitor) callOpFunction(returnType sym.Type, tokenType int, leftObject, rightObject interface{}) interface{} {
	function := v.at.GetOpFunction(returnType, tokenType)
	if function != nil {
		funcObject := stack.NewFuncObject(function)
		opParams := []interface{}{leftObject, rightObject}
		return v.executeFunctionCall(funcObject, opParams)
	} else {
		// todo crossoverJie 运行时错误 没有实现运算符重载
	}
	return nil
}

// VisitFunctionCall 函数调用
func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	var ret interface{}
	name := ctx.IDENTIFIER().GetText()
	// todo crossoverJie 内置函数校验
	if name == "println" {
		v.println(ctx)
		return ret
	} else if name == "assertEqual" {
		v.assertEqual(ctx)
		return ret
	} else if name == "append" {
		return v.append(ctx)
	} else if name == "len" {
		return v.len(ctx)
	} else if name == "hash" {
		return v.hash(ctx)
	} else if name == "JSON" {
		return v.JSON(ctx)
	} else if name == "JSONGet" {
		return v.JSONGet(ctx)
	} else if name == "httpHandle" {
		return v.httpHandle(ctx)
	} else if name == "httpRun" {
		return v.httpRun(ctx)
	} else if name == "FprintfJSON" {
		v.fprintfJSON(ctx)
	} else if name == "FprintfHTML" {
		v.fprintfHTML(ctx)
	} else if name == "getCurrentTime" {
		return v.getCurrentTime(ctx)
	} else if name == "QueryPath" {
		return v.queryPath(ctx)
	} else if name == "FormValue" {
		return v.formValue(ctx)
	} else if name == "getOSArgs" {
		return v.getOSArgs(ctx)
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

	//如果是 returnObject，需要把真正的数据取出
	switch ret.(type) {
	case *ReturnObject:
		ret = ret.(*ReturnObject).GetReturnObject()
	}

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
	if ctx.String_() != nil {
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
	ret := v.VisitParExpression(ctx.ParExpression().(*parser.ParExpressionContext))
	var condition bool
	switch ret.(type) {
	case bool:
		condition = ret.(bool)
	case *LeftValue:
		condition = ret.(*LeftValue).GetValue().(bool)
	}
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

func (v *Visitor) VisitStmWhile(ctx *parser.StmWhileContext) interface{} {
	scope := v.at.GetNode2Scope()[ctx]
	v.pushStack(stack.NewBlockScopeFrame(scope))
	parExpressionCtx := ctx.ParExpression().(*parser.ParExpressionContext)
	for {
		condition := true
		value := v.VisitExpr(parExpressionCtx.Expr().(*parser.ExprContext))
		switch value.(type) {
		case *LeftValue:
			c, ok := value.(*LeftValue).GetValue().(bool)
			if ok {
				condition = c
			} else {
				// todo crossoverJie 运行报错
			}
		case bool:
			condition = value.(bool)
		default:
			// todo crossoverJie 运行报错
		}

		if !condition {
			break
		}

		ret := v.Visit(ctx.Statement())
		// break
		_, b := ret.(*BreakObject)
		if b {
			break
		}

		// return
		_, r := ret.(*ReturnObject)
		if r {
			break
		}
	}

	v.popStack()
	return nil
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

			// break
			_, b := ret.(*BreakObject)
			if b {
				break
			}

			// return
			_, r := ret.(*ReturnObject)
			if r {
				break
			}

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
	var ret interface{}
	if ctx.Expr() != nil {
		ret = v.Visit(ctx.Expr())
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

	}
	// 支持 return
	ret = NewReturnObject(ret)
	return ret

}

func (v *Visitor) VisitStmBreak(ctx *parser.StmBreakContext) interface{} {
	return BreakObjectInstance
}

func (v *Visitor) VisitStmContinue(ctx *parser.StmContinueContext) interface{} {
	return ContinueObjectInstance
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
