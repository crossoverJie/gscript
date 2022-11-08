package gscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
	"github.com/crossoverJie/gscript/symbol"
)

// SemanticResolver 语义分析
/**
1. return
	1.1 声明了返回值必须返回
	1.2 返回值类型必须和声明相同
2. break
	2.1 break 只能出现在循环中
3. type check
	3.1 赋值类型匹配
	3.2 四则运算表达式类型匹配
4. 表达式运行类型检查
*/
type SemanticResolver struct {
	parser.BaseGScriptListener
	at *resolver.AnnotatedTree
}

func NewSemanticResolver(at *resolver.AnnotatedTree) *SemanticResolver {
	return &SemanticResolver{at: at}
}

func (s *SemanticResolver) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	function := s.at.GetNode2Scope()[ctx]
	if ctx.TypeTypeOrVoid() != nil {
		switch function.(type) {
		case *symbol.Func:
			f := function.(*symbol.Func)
			if GetInternalFunction(f.GetName()) != nil {
				// skip internal function
				return
			}
			returnType := f.GetReturnType()
			_, isDeclareFunction := returnType.(*symbol.DeclareFunctionType)
			r := s.checkReturn(ctx)
			if returnType != nil && r == false && isDeclareFunction == false {
				// 函数声明了返回值，但却没有返回返回值时，同时这个函数声明的返回值是"函数类型"时，则不需要校验
				s.at.Log(ctx, fmt.Sprintf("miss return at end of function: %s", f.GetName()))
			}
		}
	}
}

// 递归查询下级是否包含 return statement
func (s *SemanticResolver) checkReturn(ctx antlr.Tree) bool {
	for _, tree := range ctx.GetChildren() {
		switch tree.(type) {
		case *parser.BlockStmContext:
			stm := tree.(*parser.BlockStmContext)
			_, ok := stm.Statement().(*parser.StmReturnContext)
			if ok {
				return true
			}
		default:
			if s.checkReturn(tree) {
				return true
			}

		}
	}
	return false
}

func (s *SemanticResolver) ExitStmReturn(ctx *parser.StmReturnContext) {
	if ctx.Expr() != nil {
		exprContext := ctx.Expr().(*parser.ExprContext)
		// 当前返回的类型
		t := s.at.GetTypeOfNode()[exprContext]
		if t != nil {
			declarationCtx := s.findFunctionDeclarationCtx(ctx)
			if declarationCtx == nil {
				return
			}
			function := s.at.GetNode2Scope()[declarationCtx]
			if function == nil {
				return
			}
			switch function.(type) {
			case *symbol.Func:
				f := function.(*symbol.Func)
				if GetInternalFunction(f.GetName()) != nil {
					// skip internal function
					return
				}
				if !f.GetReturnType().IsType(t) {
					declareFunctionType, ok := t.(*symbol.DeclareFunctionType)
					if ok {
						// 函数返回值是闭包调用时：
						if !f.GetReturnType().IsType(declareFunctionType.GetReturnType()) {
							s.at.Log(ctx, fmt.Sprintf("cannot use type %s as type %s in return closure statment", declareFunctionType.GetReturnType().GetName(), f.GetReturnType().GetName()))
						}
					} else {
						// 函数返回类型与当前 return 类型不符时候
						s.at.Log(ctx, fmt.Sprintf("cannot use type %s as type %s in return statment", t.GetName(), f.GetReturnType().GetName()))
					}
				}
			}
		}
	}
}

// 查询函数声明 ctx
func (s *SemanticResolver) findFunctionDeclarationCtx(ctx antlr.Tree) *parser.FunctionDeclarationContext {
	if ctx == nil || ctx.GetParent() == nil {
		return nil
	}
	parent := ctx.GetParent()
	functionDeclareCtx, ok := parent.(*parser.FunctionDeclarationContext)
	if ok {
		return functionDeclareCtx
	}
	return s.findFunctionDeclarationCtx(parent.GetParent())

}

// ExitVariableDeclarator 赋值类型校验
func (s *SemanticResolver) ExitVariableDeclarator(ctx *parser.VariableDeclaratorContext) {
	if ctx.VariableDeclaratorId() == nil {
		return
	}
	sym := s.at.GetSymbolOfNode()[ctx.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext)]
	variable, ok := sym.(*symbol.Variable)
	if !ok {
		return
	}

	if ctx.VariableInitializer() == nil || ctx.VariableInitializer().(*parser.VariableInitializerContext).Expr() == nil {
		return
	}
	exprContext := ctx.VariableInitializer().(*parser.VariableInitializerContext).Expr().(*parser.ExprContext)
	exprType := s.at.GetTypeOfNode()[exprContext]
	if exprType == nil {
		return
	}

	// 数组赋值校验
	if variable.IsArray() && variable.GetType() != symbol.Any {
		if !exprType.IsArray() {
			s.at.Log(ctx, fmt.Sprintf("cannot use %s as type %s[]", variable.GetName(), variable.GetType().GetName()))
			return
		}
	}

	declareFunctionType, ok := exprType.(*symbol.DeclareFunctionType)
	if ok {
		// 变量为闭包变量 func int(int) f2 = f1();
		declareFuncType, ok := variable.GetType().(*symbol.DeclareFunctionType)
		if ok {
			if declareFuncType.GetReturnType() != nil && !declareFuncType.GetReturnType().IsType(declareFunctionType.GetReturnType()) {
				s.at.Log(ctx, fmt.Sprintf("cannot use type %s as type %s in assignment closure statment", declareFunctionType.GetReturnType().GetName(), declareFuncType.GetName()))
				return
			}
		} else {
			// = 闭包：
			if !variable.GetType().IsType(declareFunctionType.GetReturnType()) {
				s.at.Log(ctx, fmt.Sprintf("cannot use type %s as type %s in assignment closure statment", declareFunctionType.GetReturnType().GetName(), variable.GetType().GetName()))
				return
			}
		}

		return
	}

	functionType, ok := variable.GetType().(*symbol.DeclareFunctionType)
	if ok {
		// 	func int(int) f2 = f1(); f2 作为变量
		if functionType.GetReturnType() != nil && !functionType.GetReturnType().IsType(exprType) {
			s.at.Log(ctx, fmt.Sprintf("variable %s type error, %s not match %s in assignment closure statment", variable.GetName(), variable.GetType().GetName(), exprType.GetName()))
		}
	} else {
		if !variable.GetType().IsType(exprType) {
			s.at.Log(ctx, fmt.Sprintf("variable %s type error, %s not match %s", variable.GetName(), variable.GetType().GetName(), exprType.GetName()))
		}
	}

}

// ExitExpr 表达式运算类型检查
func (s *SemanticResolver) ExitExpr(ctx *parser.ExprContext) {
	if ctx.GetBop() != nil && len(ctx.AllExpr()) >= 2 {
		//推导出来的该节点类型
		deriveType := s.at.GetTypeOfNode()[ctx]
		type1 := s.at.GetTypeOfNode()[ctx.Expr(0)]
		type2 := s.at.GetTypeOfNode()[ctx.Expr(1)]
		if type1 == nil || type2 == nil {
			// 没有定义变量的情况 a+2;
			return
		}
		switch ctx.GetBop().GetTokenType() {
		case parser.GScriptParserMULT:
			//deriveType := symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, type1, ctx.GetBop().GetTokenType(), "*")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v * %v", type1.GetName(), type2.GetName()))
			}

		case parser.GScriptParserDIV:
			if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, type1, ctx.GetBop().GetTokenType(), "/")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v / %v", type1.GetName(), type2.GetName()))
			}

		case parser.GScriptParserPLUS:
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, type1, ctx.GetBop().GetTokenType(), "+")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v + %v", type1.GetName(), type2.GetName()))

			}
		case parser.GScriptParserSUB:
			if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, type1, ctx.GetBop().GetTokenType(), "-")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v - %v", type1.GetName(), type2.GetName()))

			}

		case parser.GScriptParserMOD:
			if deriveType == symbol.Int {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, type1, ctx.GetBop().GetTokenType(), "%")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v mod %v", type1.GetName(), type2.GetName()))

			}
		case parser.GScriptParserGT:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
				// 字符串比较永远都是 false
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), ">")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v > %v", type1.GetName(), type2.GetName()))
			}
		case parser.GScriptParserLT:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), "<")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v < %v", type1.GetName(), type2.GetName()))
			}
		case parser.GScriptParserGE:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), ">=")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v >= %v", type1.GetName(), type2.GetName()))

			}
		case parser.GScriptParserLE:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), "<=")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v <= %v", type1.GetName(), type2.GetName()))
			}
		case parser.GScriptParserEQUAL:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if deriveType == symbol.Nil {
			} else if deriveType == symbol.Any {
				// 两个 any 值进行比较
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), "==")
			}
		case parser.GScriptParserNOTEQUAL:
			deriveType = symbol.GetUpperType(ctx, type1, type2)
			if deriveType == symbol.String {
			} else if deriveType == symbol.Int {
			} else if deriveType == symbol.Byte {
			} else if deriveType == symbol.Float {
			} else if deriveType == symbol.Nil {
			} else if deriveType == symbol.Any {
			} else if type1.IsType(type2) {
				// 两个参数类型相同，执行运算符重载
				s.checkOpFunction(ctx, symbol.Bool, ctx.GetBop().GetTokenType(), "!=")
			} else {
				s.at.Log(ctx, fmt.Sprintf("invalid operation: %v != %v", type1.GetName(), type2.GetName()))

			}
		}
	}
}

func (s *SemanticResolver) checkOpFunction(ctx antlr.ParserRuleContext, returnType symbol.Type, tokenType int, op string) {
	// 基本类型不需要查询运算符重载函数
	if returnType == symbol.Any || returnType.IsType(symbol.Int) || returnType.IsType(symbol.String) || returnType.IsType(symbol.Float) || returnType.IsType(symbol.Bool) || returnType.IsType(symbol.Byte) || returnType.IsType(symbol.Nil) {
		return
	}
	function := s.at.GetOpFunction(returnType, tokenType)
	if function == nil {
		s.at.Log(ctx, fmt.Sprintf("no match to operator overloading %s function, return type: %s", op, returnType.GetName()))
	}
}
