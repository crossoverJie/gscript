package resolver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/container"
	"github.com/crossoverJie/gscript/symbol"
)

// ClosureResolver 闭包分析，为函数生成闭包变量
type ClosureResolver struct {
	at *AnnotatedTree
}

func NewClosureResolver(at *AnnotatedTree) *ClosureResolver {
	return &ClosureResolver{at: at}
}

// Analyze 计算闭包变量：所有声明变量-内部声明变量
func (c *ClosureResolver) Analyze() {
	// todo crossoverJie type_scope_resolver.go 完成后改为并发扫描
	for _, t := range c.at.GetTypes() {
		function, ok := t.(*symbol.Func)
		if ok && !function.IsMethod() {
			allVariable := c.allVariable(function)
			scopeVariable := c.currentScopeVariable(function)
			allVariable.RemoveAll(scopeVariable)
			function.SetClosureVar(allVariable)
		}
	}
}

// 所有声明的变量，函数中存在内部函数的情况，内部函数也算
func (c *ClosureResolver) allVariable(scope symbol.Scope) container.Set {
	set := container.NewSet()
	for ctx, s := range c.at.GetSymbolOfNode() {
		variable, ok := s.(*symbol.Variable)
		if ok && c.isParentCtx(scope.GetCtx(), ctx) {
			set.Add(variable)
		}
	}
	return set
}

// 当前函数所在的 scope 中的声明的变量
func (c *ClosureResolver) currentScopeVariable(scope symbol.Scope) container.Set {
	set := container.NewSet()
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *symbol.Variable:
			set.Add(s)
		case symbol.Scope:
			set.AddAll(c.currentScopeVariable(s.(symbol.Scope)))
		}
	}

	return set
}

// c1 is parent of c2
func (c *ClosureResolver) isParentCtx(c1, c2 antlr.ParserRuleContext) bool {
	if c2.GetParent() == nil {
		return false
	} else if c2.GetParent() == c1 {
		return true
	} else {
		return c.isParentCtx(c1, c2.GetParent().(antlr.ParserRuleContext))
	}
}
