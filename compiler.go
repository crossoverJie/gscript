package gscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/internal"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
)

var Args []string

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

// CompilerWithoutNative 不包含标准库运行
func (c *Compiler) CompilerWithoutNative(script string) interface{} {
	return c.compile(script)
}

func (c *Compiler) compile(script string) interface{} {
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

func (c *Compiler) Compiler(script string) interface{} {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r)
	//	}
	//}()
	internal := c.loadInternal()
	script = internal + script
	return c.compile(script)
}

func (c *Compiler) loadInternal() string {
	bytes, err := internal.Asset("internal/internal.gs")
	//file, err := os.ReadFile("internal/internal.gs")
	if err != nil {
		panic(err)
	}
	return string(bytes)

}
