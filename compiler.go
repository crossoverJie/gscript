package gscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
)

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compiler(script string) interface{} {
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	tree := parser.NewGScriptParser(stream).Prog()
	at := resolver.NewAnnotatedTree(tree)

	walker := antlr.NewParseTreeWalker()
	walker.Walk(resolver.NewTypeScopeResolver(at), tree)

	visitor := NewVisitor(at)
	return visitor.Visit(tree)
}
