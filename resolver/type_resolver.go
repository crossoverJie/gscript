package resolver

import (
	"fmt"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/symbol"
)

// TypeResolver 第二次扫描，记录变量，类继承、函数声明等信息到 symbolOfNode
// 设置变量类型
// 对局部变量需要边添加边解析，不然会引起变量消解错误，比如父级 scope 和 局部变量中存在名称相同的变量，会导致局部变量消解到父级中的变量。
type TypeResolver struct {
	parser.BaseGScriptListener
	at              *AnnotatedTree
	isLocalVariable bool
}

func NewTypeResolver(at *AnnotatedTree) *TypeResolver {
	return &TypeResolver{at: at}
}

func NewTypeResolverWithLocalVariable(at *AnnotatedTree) *TypeResolver {
	return &TypeResolver{at: at, isLocalVariable: true}
}

// ExitVariableDeclarators 设置变量类型
func (t *TypeResolver) ExitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) {
	// todo crossoverJie scope 是 class
	_ = t.at.FindEncloseScopeOfNode(ctx)
	if t.isLocalVariable {
		// 在 ExitTypeType 设置的类型
		symbolType := t.at.GetTypeOfNode()[ctx.TypeType()]
		for _, context := range ctx.AllVariableDeclarator() {
			s := t.at.GetSymbolOfNode()[context]
			switch s.(type) {
			case *symbol.Variable:
				s.(*symbol.Variable).SetType(symbolType)
			}
		}
	}
}

func (t *TypeResolver) EnterVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) {
	id := ctx.IDENTIFIER().GetText()
	scope := t.at.FindEncloseScopeOfNode(ctx)

	// todo crossoverJie scope is class
	_, ok := ctx.GetParent().(*parser.FormalParameterContext)
	if t.isLocalVariable || ok {
		// 本地变量/ class/ parent.FormalParameterContext(.g4 105)函数中的参数才写入
		variable := symbol.NewVariable(ctx, id, scope)

		if symbol.GetVariable(scope, id) != nil {
			panic(fmt.Sprintf("repeat variable %s, at %d", id, ctx.GetStart().GetLine()))
		}

		scope.AddSymbol(variable)
		t.at.PutSymbolOfNode(ctx, variable)
	}
}

// ExitFunctionDeclaration 函数声明，设置函数的返回类型
func (t *TypeResolver) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	// TypeScopeResolver.EnterFunctionDeclaration() 写入的数据
	function := t.at.GetNode2Scope()[ctx]
	if ctx.TypeTypeOrVoid() != nil {
		// ExitTypeTypeOrVoid 设置的返回值类型
		returnType := t.at.GetTypeOfNode()[ctx.TypeTypeOrVoid()]
		switch function.(type) {
		case *symbol.Func:
			function.(*symbol.Func).SetReturnType(returnType)
		}
	}

	// todo crossoverJie func 查询是否重复

}

// ExitFormalParameter 函数入参，设置函数入参类型
func (t *TypeResolver) ExitFormalParameter(ctx *parser.FormalParameterContext) {
	// 获取在 typeType 中声明的类型
	symbolType := t.at.GetTypeOfNode()[ctx.TypeType()]

	// 获取在 EnterVariableDeclaratorId 中设置的变量
	variable := t.at.GetSymbolOfNode()[ctx.VariableDeclaratorId()]
	switch variable.(type) {
	case *symbol.Variable:
		// 为变量设置类型
		v := variable.(*symbol.Variable)
		v.SetType(symbolType)

		scope := t.at.FindEncloseScopeOfNode(ctx)
		switch scope.(type) {
		case *symbol.Func:
			// 为函数设置入参
			scope.(*symbol.Func).AppendParameter(v)
		}
	}

}

func (t *TypeResolver) ExitTypeTypeOrVoid(ctx *parser.TypeTypeOrVoidContext) {
	if ctx.TypeType() != nil {
		// ExitTypeType 写入的类型
		typeType := t.at.GetTypeOfNode()[ctx.TypeType()]
		t.at.PutTypeOfNode(ctx, typeType)
	} else if ctx.VOID() != nil {
		t.at.PutTypeOfNode(ctx, symbol.Void)
	}
}

// ExitTypeType 在 ExitPrimitiveType 之后记录当前 ctx 的基本类型
func (t *TypeResolver) ExitTypeType(ctx *parser.TypeTypeContext) {
	// todo crossoverJie func class 的 type 还未处理
	if ctx.PrimitiveType() != nil {
		symbolType := t.at.GetTypeOfNode()[ctx.PrimitiveType()]
		t.at.PutTypeOfNode(ctx, symbolType)
	}
}

// ExitPrimitiveType 记录当前 ctx 的基本类型
func (t *TypeResolver) ExitPrimitiveType(ctx *parser.PrimitiveTypeContext) {
	var symbolType symbol.Type
	if ctx.INT() != nil {
		symbolType = symbol.Int
	}
	if ctx.STRING() != nil {
		symbolType = symbol.String
	}
	if ctx.FLOAT() != nil {
		symbolType = symbol.Float
	}
	if ctx.BOOLEAN() != nil {
		symbolType = symbol.Bool
	}
	t.at.PutTypeOfNode(ctx, symbolType)
}
