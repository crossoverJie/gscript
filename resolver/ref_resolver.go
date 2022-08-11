package resolver

import (
	"fmt"
	"github.com/crossoverJie/gscript/parser"
)

type RefResolver struct {
	parser.BaseGScriptListener
	at *AnnotatedTree
}

func NewRefResolver(at *AnnotatedTree) *RefResolver {
	return &RefResolver{at: at}
}

func (s *RefResolver) ExitIdentifierPrimary(ctx *parser.IdentifierPrimaryContext) {
	idName := ctx.IDENTIFIER().GetText()
	scope := s.at.FindEncloseScopeOfNode(ctx)
	variable := s.at.FindVariable(scope, idName)
	// todo crossoverJie 区分返回的是函数
	if variable == nil {
		// todo crossoverJie 完善编译报错信息
		panic(fmt.Sprintf("unknown variable %s, at %d", idName, ctx.GetStart().GetLine()))
	}

	s.at.PutSymbolOfNode(ctx, variable)

	// todo crossoverJie 类型推导
}
