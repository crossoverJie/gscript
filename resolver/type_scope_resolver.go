package resolver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
)

type TypeScopeResolver struct {
	parser.BaseGScriptListener
	stack stack.Stack
	at    *AnnotatedTree
}

func NewTypeScopeResolver(at *AnnotatedTree) *TypeScopeResolver {
	return &TypeScopeResolver{at: at}
}

func (t *TypeScopeResolver) pushScope(ctx antlr.ParserRuleContext, scope symbol.Scope) {
	t.at.PutNode2Scope(ctx, scope)
	scope.SetCtx(ctx)
	t.stack.Push(scope)
}

func (t *TypeScopeResolver) popScope() {
	t.stack.Pop()
}

func (t *TypeScopeResolver) currentScope() symbol.Scope {
	if !t.stack.IsEmpty() {
		return t.stack.Peek().(symbol.Scope)
	}
	return nil
}

func (t *TypeScopeResolver) EnterProg(ctx *parser.ProgContext) {
	scope := symbol.NewBlockScope(ctx, "prog", t.currentScope())
	t.pushScope(ctx, scope)
}

func (t *TypeScopeResolver) ExitProg(ctx *parser.ProgContext) {
	t.popScope()
}
