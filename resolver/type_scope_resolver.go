package resolver

import (
	"fmt"
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

	// 函数类型则不需要创建额外的 scope
	_, ok := ctx.GetParent().(*parser.FunctionBodyContext)
	if !ok {
		scope := symbol.NewBlockScope(ctx, "block", t.currentScope())
		t.currentScope().AddSymbol(scope)
		t.pushScope(ctx, scope)
	}
}

// ExitBlock is called when production block is exited.
func (t *TypeScopeResolver) ExitBlock(ctx *parser.BlockContext) {
	// 函数类型则不需要创建额外的 scope
	_, ok := ctx.GetParent().(*parser.FunctionBodyContext)
	if !ok {
		t.popScope()
	}
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

// EnterStmWhile is called when production StmWhile is entered.
func (t *TypeScopeResolver) EnterStmWhile(ctx *parser.StmWhileContext) {
	scope := symbol.NewBlockScope(ctx, "while", t.currentScope())
	t.currentScope().AddSymbol(scope)
	t.pushScope(ctx, scope)
}

// ExitStmWhile is called when production StmWhile is exited.
func (t *TypeScopeResolver) ExitStmWhile(ctx *parser.StmWhileContext) {
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

	// 保存当前 ctx 的函数，用于递归函数判断
	t.at.PutFunction2Scope(ctx, newFunc)

	t.pushScope(ctx, newFunc)
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (t *TypeScopeResolver) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	t.popScope()
}

func (t *TypeScopeResolver) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	className := ctx.IDENTIFIER().GetText()
	findClass := t.at.FindClass(t.currentScope(), className)
	if findClass != nil {
		t.at.Log(ctx, fmt.Sprintf("class %s redeclared in this block", className))
	}

	class := symbol.NewClass(ctx, className)
	t.currentScope().AddSymbol(class)
	t.at.AppendType(class)

	t.pushScope(ctx, class)
}

func (t *TypeScopeResolver) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	t.popScope()
}

// EnterOperatorOverloading is called when production operatorOverloading is entered.
func (t *TypeScopeResolver) EnterOperatorOverloading(ctx *parser.OperatorOverloadingContext) {
	var opOverload *symbol.OpOverload
	function, ok := t.currentScope().(*symbol.Func)
	if !ok || function.GetName() != symbol.OperatorName {
		return
	}
	if ctx.PLUS() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.PLUS().GetSymbol().GetTokenType())
	} else if ctx.SUB() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.SUB().GetSymbol().GetTokenType())
	} else if ctx.MULT() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.MULT().GetSymbol().GetTokenType())
	} else if ctx.DIV() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.DIV().GetSymbol().GetTokenType())
	} else if ctx.EQUAL() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.EQUAL().GetSymbol().GetTokenType())
	} else if ctx.NOTEQUAL() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.NOTEQUAL().GetSymbol().GetTokenType())
	} else if ctx.GT() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.GT().GetSymbol().GetTokenType())
	} else if ctx.LT() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.LT().GetSymbol().GetTokenType())
	} else if ctx.GE() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.GE().GetSymbol().GetTokenType())
	} else if ctx.LE() != nil {
		opOverload = symbol.NewOpOverload(function, ctx.LE().GetSymbol().GetTokenType())
	}
	if opOverload != nil {
		t.at.AppendOpOverload(opOverload)
	}
}

// ExitOperatorOverloading is called when production operatorOverloading is exited.
func (t *TypeScopeResolver) ExitOperatorOverloading(ctx *parser.OperatorOverloadingContext) {
	//t.popScope()
}
