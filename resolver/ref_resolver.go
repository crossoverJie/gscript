package resolver

import "github.com/crossoverJie/gscript/parser"

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
	s.at.PutSymbolOfNode(ctx, variable)

	// todo crossoverJie 类型推导
}
