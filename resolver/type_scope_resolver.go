package resolver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
)

// TypeScopeResolver 第一次扫描，记录所有的类型：类、函数、scope 到 node2Scope 中。
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

// EnterBlock is called when production block is entered.
func (t *TypeScopeResolver) EnterBlock(ctx *parser.BlockContext) {

	// todo crossoverJie 函数类型则不需要创建额外的 scope
	scope := symbol.NewBlockScope(ctx, "block", t.currentScope())
	t.currentScope().AddSymbol(scope)
	t.pushScope(ctx, scope)
}

// ExitBlock is called when production block is exited.
func (t *TypeScopeResolver) ExitBlock(ctx *parser.BlockContext) {
	// todo crossoverJie 函数类型则不需要创建额外的 scope
	t.popScope()
}

// EnterStmFor is called when production StmFor is entered.
func (t *TypeScopeResolver) EnterStmFor(ctx *parser.StmForContext) {
	scope := symbol.NewBlockScope(ctx, "for", t.currentScope())
	t.currentScope().AddSymbol(scope)
	t.pushScope(ctx, scope)
}

// ExitStmFor is called when production StmFor is exited.
func (t *TypeScopeResolver) ExitStmFor(ctx *parser.StmForContext) {
	t.popScope()
}

// EnterFunctionDeclaration 加入一个 func 的 scope
func (t *TypeScopeResolver) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	//scope := symbol.NewBlockScope(ctx, "func", t.currentScope())
	//t.currentScope().AddSymbol(scope)

	name := ctx.IDENTIFIER().GetText()
	newFunc := symbol.NewFunc(ctx, name, t.currentScope())
	t.currentScope().AddSymbol(newFunc)

	// add type
	t.at.AppendType(newFunc)

	t.pushScope(ctx, newFunc)
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (t *TypeScopeResolver) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	t.popScope()
}
