package gscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/internal"
	"github.com/crossoverJie/gscript/log"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/resolver"
	"os"
	"strings"
)

var Args []string

const (
	RuntimeError = "RuntimeError"
)

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

// CompilerWithoutNative 不包含标准库运行
func (c *Compiler) CompilerWithoutNative(script string) interface{} {
	return c.compile(script)
}

// GetCompileInfo 获取编译的 AST 以及符号表
func (c *Compiler) GetCompileInfo(script string, isAST bool) string {
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	scriptParser := parser.NewGScriptParser(stream)
	tree := scriptParser.Prog()
	at := resolver.NewAnnotatedTree(tree, scriptParser)

	walker := antlr.NewParseTreeWalker()
	// 识别所有的类型、函数、scope
	walker.Walk(resolver.NewTypeScopeResolver(at), tree)

	// 变量、类型解析，所有使用到 typeType 的地方
	walker.Walk(resolver.NewTypeResolver(at), tree)

	if isAST {
		return at.DumpAST()
	} else {
		return at.DumpSymbol()
	}
}

func (c *Compiler) compile(script string) interface{} {
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	scriptParser := parser.NewGScriptParser(stream)
	tree := scriptParser.Prog()
	at := resolver.NewAnnotatedTree(tree, scriptParser)

	walker := antlr.NewParseTreeWalker()
	// 识别所有的类型、函数、scope
	walker.Walk(resolver.NewTypeScopeResolver(at), tree)

	// 变量、类型解析，所有使用到 typeType 的地方
	walker.Walk(resolver.NewTypeResolver(at), tree)

	// 消解变量、函数的引用
	walker.Walk(resolver.NewRefResolver(at), tree)

	// 语义检查
	walker.Walk(NewSemanticResolver(at), tree)

	// 闭包分析
	resolver.NewClosureResolver(at).Analyze()

	if at.IsCompileFail() {
		return nil
	}
	visitor := NewVisitor(at)
	return visitor.Visit(tree)
}

func (c *Compiler) Compiler(script string) interface{} {
	runtimeError := os.Getenv(RuntimeError)
	if runtimeError != "" {
		defer func() {
			if r := recover(); r != nil {
				switch x := r.(type) {
				case *log.Log:
					fmt.Printf("runtime error: %s \n", x.String())
				default:
					panic(x)
				}
			}
		}()
	}
	native := c.loadInternal()
	script = native + script
	return c.compile(script)
}

func (c *Compiler) loadInternal() string {
	bytes, err := internal.Asset("internal/internal.gs")
	//file, err := os.ReadFile("internal/internal.gs")
	if err != nil {
		panic(err)
	}
	v := string(bytes)
	log.NativeLine = strings.Count(v, "\n") + 1
	return v

}
