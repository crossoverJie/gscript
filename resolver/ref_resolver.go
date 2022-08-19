package resolver

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/symbol"
)

// RefResolver 引用消解和类型推断
// 函数调用，消解出是哪个 function，以及返回值类型
type RefResolver struct {
	parser.BaseGScriptListener
	at                    *AnnotatedTree
	localVariableResolver *TypeResolver
	typeResolverWalker    *antlr.ParseTreeWalker
}

func NewRefResolver(at *AnnotatedTree) *RefResolver {
	return &RefResolver{
		at:                    at,
		localVariableResolver: NewTypeResolverWithLocalVariable(at),
		typeResolverWalker:    antlr.NewParseTreeWalker(),
	}
}

// EnterVariableDeclarators 本地变量必须是添加和解析同时进行
func (s *RefResolver) EnterVariableDeclarators(ctx *parser.VariableDeclaratorsContext) {
	scope := s.at.FindEncloseScopeOfNode(ctx)
	_, isBlock := scope.(*symbol.BlockScope)
	_, isFunc := scope.(*symbol.Func)
	if isBlock || isFunc {
		s.typeResolverWalker.Walk(s.localVariableResolver, ctx)
	}
}

//func (s *RefResolver) ExitIdentifierPrimary(ctx *parser.IdentifierPrimaryContext) {
//	idName := ctx.IDENTIFIER().GetText()
//	scope := s.at.FindEncloseScopeOfNode(ctx)
//	variable := s.at.FindVariable(scope, idName)
//	if variable == nil {
//		// todo crossoverJie 区分返回的是函数
//
//		// todo crossoverJie 完善编译报错信息
//		line := ctx.GetStart().GetLine()
//		column := ctx.GetStart().GetColumn()
//		panic(fmt.Sprintf("unknown variable %s, at %d and column:%d", idName, line, column))
//	}
//
//	s.at.PutSymbolOfNode(ctx, variable)
//
//	// 写入类型1
//	s.at.PutTypeOfNode(ctx.GetParent().(*parser.ExprContext), variable.GetType())
//}

//func (s *RefResolver) ExitLiterPrimary(ctx *parser.LiterPrimaryContext) {
//	symbolType := s.at.GetTypeOfNode()[ctx]
//	s.at.PutTypeOfNode(ctx, symbolType)
//}
//
//func (s *RefResolver) ExitExprPrimary(ctx *parser.ExprPrimaryContext) {
//	symbolType := s.at.GetTypeOfNode()[ctx]
//	s.at.PutTypeOfNode(ctx, symbolType)
//}

func (s *RefResolver) ExitPrimary(ctx *parser.PrimaryContext) {
	// todo crossoverJie type 提出来
	scope := s.at.FindEncloseScopeOfNode(ctx)
	var symbolType symbol.Type
	if ctx.Expr() != nil {
		symbolType = s.at.GetTypeOfNode()[ctx.Expr()]
	}
	if ctx.Literal() != nil {
		symbolType = s.at.GetTypeOfNode()[ctx.Literal()]
	}
	if ctx.IDENTIFIER() != nil {
		idName := ctx.IDENTIFIER().GetText()
		variable := s.at.FindVariable(scope, idName)
		if variable == nil {
			// todo crossoverJie 区分返回的是函数

			// todo crossoverJie 完善编译报错信息
			line := ctx.GetStart().GetLine()
			column := ctx.GetStart().GetColumn()
			panic(fmt.Sprintf("unknown variable %s, at %d and column:%d", idName, line, column))
		} else {
			s.at.PutSymbolOfNode(ctx, variable)
			symbolType = variable.GetType()
		}
	}

	s.at.PutTypeOfNode(ctx, symbolType)
}

// ExitFunctionCall 设置当前 scope 中的函数以及函数返回值
func (s *RefResolver) ExitFunctionCall(ctx *parser.FunctionCallContext) {

	// todo crossoverJie 处理内置函数
	name := ctx.IDENTIFIER().GetText()

	paramTypes := s.getParamTypes(ctx)

	scope := s.at.FindEncloseScopeOfNode(ctx)
	found := false

	// todo crossoverJie . 符号级联调用

	if !found {
		function := s.at.FindFunction(scope, name, paramTypes)
		if function != nil {
			found = true
			s.at.PutSymbolOfNode(ctx, function)
			s.at.PutTypeOfNode(ctx, function.GetReturnType())
		}
	}

	// 查找是否是类的构造函数
	if !found {
		// 构造函数没有返回值
		class := s.at.FindClass(scope, name)
		if class != nil {
			// 查找显式的构造函数
			function := class.GetConstructorFunc(paramTypes)
			if function != nil {
				found = true
				s.at.PutSymbolOfNode(ctx, function)
			} else if len(paramTypes) == 0 {
				// 查找默认的构造函数
				found = true
				s.at.PutSymbolOfNode(ctx, class.GetDefaultConstructorFunc())
			} else {
				// todo crossoverJie 完善编译报错信息
			}
		}
	}

	// todo crossoverJie 查找是否为函数变量

}

func (s *RefResolver) ExitExpr(ctx *parser.ExprContext) {
	if ctx.GetBop() != nil && ctx.GetBop().GetTokenType() == parser.GScriptParserDOT {
		// 获取到 xx.age 中写入的 symbol
		sym := s.at.GetSymbolOfNode()[ctx.Expr(0)]
		switch sym.(type) {
		case *symbol.Variable:
			variable := sym.(*symbol.Variable)
			switch variable.GetType().(type) {
			case *symbol.Class:
				// 在 class 中通过变量名称查找变量
				name := ctx.IDENTIFIER().GetText()
				class := variable.GetType().(*symbol.Class)
				findVariable := s.at.FindVariable(class, name)
				if findVariable != nil {
					s.at.PutSymbolOfNode(ctx, findVariable)
				} else {
					// todo crossoverJie 编译报错信息
				}
			}
		}
	} else if ctx.Primary() != nil {
		// Person xx=Person();
		// xx.age 中的 xx
		sym := s.at.GetSymbolOfNode()[ctx.Primary()]
		s.at.PutSymbolOfNode(ctx, sym)
	}
	if ctx.Primary() != nil {
		symbolType := s.at.GetTypeOfNode()[ctx.Primary()]
		s.at.PutTypeOfNode(ctx, symbolType)
	} else if ctx.FunctionCall() != nil {
		// 获取方法返回值类型
		symbolType := s.at.GetTypeOfNode()[ctx.FunctionCall()]
		// 设置方法返回值类型
		s.at.PutTypeOfNode(ctx, symbolType)
	}
}

//func (s *RefResolver) ExitPrimaryExpr(ctx *parser.PrimaryExprContext) {
//	symbolType := s.at.GetTypeOfNode()[ctx]
//	s.at.PutTypeOfNode(ctx, symbolType)
//}

//func (s *RefResolver) ExitFuncCallExpr(ctx *parser.FuncCallExprContext) {
//	symbolType := s.at.GetTypeOfNode()[ctx]
//	s.at.PutTypeOfNode(ctx, symbolType)
//}

// 查询函数的参数列表
func (s *RefResolver) getParamTypes(ctx *parser.FunctionCallContext) []symbol.Type {
	var paraTypes []symbol.Type
	if ctx.ExpressionList() == nil {
		return paraTypes
	}
	for _, context := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpr() {
		symbolType := s.at.GetTypeOfNode()[context.(*parser.ExprContext)]
		paraTypes = append(paraTypes, symbolType)
	}
	return paraTypes
}

//func (s *RefResolver) ExitInt(ctx *parser.IntContext) {
//	s.at.PutTypeOfNode(ctx.GetParent().(*parser.PrimaryContext), symbol.Int)
//}

func (s *RefResolver) ExitLiteral(ctx *parser.LiteralContext) {
	if ctx.DECIMAL_LITERAL() != nil {
		// 设置标识符类型
		s.at.PutTypeOfNode(ctx, symbol.Int)
	}
}
