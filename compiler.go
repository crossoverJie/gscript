package gscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
	"log"
)

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compiler(script string) interface{} {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	tree := parser.NewGScriptParser(stream).Prog()
	at := resolver.NewAnnotatedTree(tree)

	walker := antlr.NewParseTreeWalker()
	// 识别所有的类型、函数、scope
	walker.Walk(resolver.NewTypeScopeResolver(at), tree)

	// 变量、类型解析，所有使用到 typeType 的地方
	walker.Walk(resolver.NewTypeResolver(at), tree)

	// 消解变量、函数的引用
	walker.Walk(resolver.NewRefResolver(at), tree)

	// 闭包分析
	resolver.NewClosureResolver(at).Analyze()

	visitor := NewVisitor(at)
	return visitor.Visit(tree)
}
