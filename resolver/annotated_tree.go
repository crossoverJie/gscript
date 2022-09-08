package resolver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/symbol"
)

type AnnotatedTree struct {
	ast antlr.ParseTree
	// ctx 中存放的所有作用域
	node2Scope map[antlr.ParserRuleContext]symbol.Scope

	// ctx中存放的所有符号
	symbolOfNode map[antlr.ParserRuleContext]symbol.Symbol

	// ctx 中存放的所有类型
	typeOfNode map[antlr.ParserRuleContext]symbol.Type

	// 所有的 class，function 的 type
	types []symbol.Type

	opOverloads []*symbol.OpOverload
}

func NewAnnotatedTree(ast antlr.ParseTree) *AnnotatedTree {
	return &AnnotatedTree{
		ast:          ast,
		symbolOfNode: make(map[antlr.ParserRuleContext]symbol.Symbol),
		typeOfNode:   make(map[antlr.ParserRuleContext]symbol.Type),
		node2Scope:   make(map[antlr.ParserRuleContext]symbol.Scope),
		types:        make([]symbol.Type, 0),
		opOverloads:  make([]*symbol.OpOverload, 0),
	}
}

func (a *AnnotatedTree) PutSymbolOfNode(ctx antlr.ParserRuleContext, symbol symbol.Symbol) {
	a.symbolOfNode[ctx] = symbol
}

func (a *AnnotatedTree) GetSymbolOfNode() map[antlr.ParserRuleContext]symbol.Symbol {
	return a.symbolOfNode
}

func (a *AnnotatedTree) PutTypeOfNode(ctx antlr.ParserRuleContext, t symbol.Type) {
	a.typeOfNode[ctx] = t
}

func (a *AnnotatedTree) GetTypeOfNode() map[antlr.ParserRuleContext]symbol.Type {
	return a.typeOfNode
}

func (a *AnnotatedTree) PutNode2Scope(ctx antlr.ParserRuleContext, symbol symbol.Scope) {
	a.node2Scope[ctx] = symbol
}

func (a *AnnotatedTree) GetNode2Scope() map[antlr.ParserRuleContext]symbol.Scope {
	return a.node2Scope
}

func (a *AnnotatedTree) AppendType(t symbol.Type) {
	a.types = append(a.types, t)
}

func (a *AnnotatedTree) GetTypes() []symbol.Type {
	return a.types
}
func (a *AnnotatedTree) AppendOpOverload(op *symbol.OpOverload) {
	a.opOverloads = append(a.opOverloads, op)
}

//
func (a *AnnotatedTree) GetOpOverloadWithReturnType(returnType symbol.Type, tokenType int) *symbol.Func {
	for _, overload := range a.opOverloads {
		isType := overload.GetFunc().GetReturnType().IsType(returnType)
		if isType && overload.GetTokenType() == tokenType {
			return overload.GetFunc()
		}
	}
	return nil
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
func (a *AnnotatedTree) FindClass(scope symbol.Scope, name string) *symbol.Class {
	class := scope.GetClass(name)
	if class == nil && scope.GetEncloseScope() != nil {
		class = a.FindClass(scope.GetEncloseScope(), name)
	}
	return class
}

// FindVariable 根据变量名称在 scope 中逐级查找变量
func (a *AnnotatedTree) FindVariable(scope symbol.Scope, name string) *symbol.Variable {
	var ret = scope.GetVariable(name)
	if ret == nil && scope.GetEncloseScope() != nil {
		ret = a.FindVariable(scope.GetEncloseScope(), name)
	}
	return ret
}

// FindFunctionVariable 查找函数类型的变量：变量中存放的值是 functionObject
func (a *AnnotatedTree) FindFunctionVariable(scope symbol.Scope, name string, paramTypes []symbol.Type) *symbol.Variable {
	var ret = scope.GetFunctionVariable(name, paramTypes)
	if ret == nil && scope.GetEncloseScope() != nil {
		ret = a.FindFunctionVariable(scope.GetEncloseScope(), name, paramTypes)
	}
	return ret
}

// FindFunction 根据方法名称+方法参数查询方法
func (a *AnnotatedTree) FindFunction(scope symbol.Scope, name string, paramTypes []symbol.Type) *symbol.Func {
	var ret = scope.GetFunction(name, paramTypes)
	if ret == nil && scope.GetEncloseScope() != nil {
		ret = a.FindFunction(scope.GetEncloseScope(), name, paramTypes)
	}
	return ret
}

// FindFunctionWithName 只通过函数名称查询
// 先查询 class 的method，再查询 function
func (a *AnnotatedTree) FindFunctionWithName(scope symbol.Scope, name string) *symbol.Func {
	var fun *symbol.Func
	class, ok := scope.(*symbol.Class)
	if ok {
		fun = a.findMethodWithName(class, name)
	} else {
		fun = a.findFunctionOnlyWithName(scope, name)
	}

	if fun == nil && scope.GetEncloseScope() != nil {
		fun = a.FindFunctionWithName(scope.GetEncloseScope(), name)
	}

	return fun
}

// 查找 class 的 method
func (a *AnnotatedTree) findMethodWithName(class *symbol.Class, name string) *symbol.Func {
	// todo crossoverJie 父类查询
	return a.findFunctionOnlyWithName(class, name)
}

func (a *AnnotatedTree) findFunctionOnlyWithName(scope symbol.Scope, name string) *symbol.Func {
	// todo crossoverJie 函数名重复的情况，应该在写入 symbol 时就行编译报错
	for _, s := range scope.GetSymbols() {
		function, ok := s.(*symbol.Func)
		if ok && function.GetName() == name {
			return function
		}
	}
	return nil
}
