package resolver

import (
	"fmt"
	"github.com/crossoverJie/gscript/parser"
)

type TypeScopeResolver struct {
	parser.BaseGScriptListener
}

func NewTypeScopeResolver() *TypeScopeResolver {
	return &TypeScopeResolver{}
}

func (t *TypeScopeResolver) EnterProg(ctx *parser.ProgContext) {
	fmt.Println("enter prog")
}
