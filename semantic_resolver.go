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
			r := s.checkReturn(ctx)
			if returnType != nil && r == false {
				s.at.Log(ctx, fmt.Sprintf("miss return at end of function: %s", f.GetName()))
			}
		}
	}
}

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
