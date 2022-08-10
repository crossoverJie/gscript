package resolver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/symbol"
)

type AnnotatedTree struct {
	ast          antlr.ParseTree
	symbolOfNode map[antlr.ParserRuleContext]symbol.Symbol
	node2Scope   map[antlr.ParserRuleContext]symbol.Scope
}

func NewAnnotatedTree(ast antlr.ParseTree) *AnnotatedTree {
	return &AnnotatedTree{
		ast:          ast,
		symbolOfNode: make(map[antlr.ParserRuleContext]symbol.Symbol),
		node2Scope:   make(map[antlr.ParserRuleContext]symbol.Scope),
	}
}

func (a *AnnotatedTree) PutSymbolOfNode(ctx antlr.ParserRuleContext, symbol symbol.Symbol) {
	a.symbolOfNode[ctx] = symbol
}

func (a *AnnotatedTree) GetSymbolOfNode() map[antlr.ParserRuleContext]symbol.Symbol {
	return a.symbolOfNode
}

func (a *AnnotatedTree) PutNode2Scope(ctx antlr.ParserRuleContext, symbol symbol.Scope) {
	a.node2Scope[ctx] = symbol
}

func (a *AnnotatedTree) GetNode2Scope() map[antlr.ParserRuleContext]symbol.Scope {
	return a.node2Scope
}

// FindEncloseScopeOfNode 查找某个 ctx 所在的 scope，逐级递归 ctx 查找
func (a *AnnotatedTree) FindEncloseScopeOfNode(ctx antlr.ParserRuleContext) symbol.Scope {
	var ret symbol.Scope
	parent := ctx.GetParent()
	if parent != nil {
		ret = a.node2Scope[parent.(antlr.ParserRuleContext)]
		if ret == nil {
			ret = a.FindEncloseScopeOfNode(parent.(antlr.ParserRuleContext))
		} else {
			return ret
		}

	}

	return ret
}

// FindVariable 根据变量名称在 scope 中逐级查找变量
func (a *AnnotatedTree) FindVariable(scope symbol.Scope, name string) *symbol.Variable {
	var ret = scope.GetVariable(name)
	if ret == nil && scope.GetEncloseScope() != nil {
		ret = a.FindVariable(scope.GetEncloseScope(), name)
	}
	return ret
}
