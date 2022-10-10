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
	scope := t.at.FindEncloseScopeOfNode(ctx)
	_, ok := scope.(*symbol.Class)
	if t.isLocalVariable || ok {
		// 在 ExitTypeType 设置的类型
		symbolType := t.at.GetTypeOfNode()[ctx.TypeType()]
		for _, context := range ctx.AllVariableDeclarator() {
			// 为变量设置类型
			s := t.at.GetSymbolOfNode()[context.(*parser.VariableDeclaratorContext).VariableDeclaratorId().(*parser.VariableDeclaratorIdContext)]
			switch s.(type) {
			case *symbol.Variable:
				s.(*symbol.Variable).SetType(symbolType)
				s.(*symbol.Variable).SetArray(symbolType.IsArray())
			}
		}
	}
}

func (t *TypeResolver) EnterVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) {
	id := ctx.IDENTIFIER().GetText()
	scope := t.at.FindEncloseScopeOfNode(ctx)
	class, isClass := scope.(*symbol.Class)

	// todo crossoverJie scope is class
	_, ok := ctx.GetParent().(*parser.FormalParameterContext)
	_, last := ctx.GetParent().(*parser.LastFormalParameterContext)
	if t.isLocalVariable || ok || isClass || last {
		// 本地变量/ class/ parent.FormalParameterContext(.g4 105)函数中的参数才写入
		variable := symbol.NewVariable(ctx, id, scope, last)

		if symbol.GetVariable(scope, id) != nil {
			panic(fmt.Sprintf("repeat variable %s, at %d", id, ctx.GetStart().GetLine()))
		}

		scope.AddSymbol(variable)
		t.at.PutSymbolOfNode(ctx, variable)

		if class != nil && class.GetName() == "HttpContext" {
			// 如果有使用 http，需要将当前的 path 变量单独存放起来，用于运行时动态获取 path。
			t.at.SetHttpPathVariable(variable)
		}
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

	switch function.(type) {
	case *symbol.Func:
		fn := function.(*symbol.Func)
		scope := t.at.FindEncloseScopeOfNode(ctx)
		found := t.at.FindFunction(scope, fn.GetName(), fn.GetParameterType())
		if found != nil && found != function && found.GetName() != symbol.OperatorName {
			t.at.Log(ctx, fmt.Sprintf("%s already declared in this block", found.GetName()))
		}
	}

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

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (t *TypeResolver) ExitLastFormalParameter(ctx *parser.LastFormalParameterContext) {
	if ctx.ELLIPSIS() == nil {
		// todo crossoverJie ?
	}
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

func (t *TypeResolver) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	// todo crossoverJie 设置父类
}

func (t *TypeResolver) ExitTypeTypeOrVoid(ctx *parser.TypeTypeOrVoidContext) {
	if ctx.TypeType() != nil {
		// ExitTypeType 写入的类型
		typeType := t.at.GetTypeOfNode()[ctx.TypeType()]
		// 方法返回值的类型
		t.at.PutTypeOfNode(ctx, typeType)
	} else if ctx.VOID() != nil {
		t.at.PutTypeOfNode(ctx, symbol.Void)
	}
}

// ExitTypeType 在 ExitPrimitiveType 之后记录当前 ctx 的基本类型
func (t *TypeResolver) ExitTypeType(ctx *parser.TypeTypeContext) {
	// todo crossoverJie class 的 type 还未处理
	var symbolType symbol.Type
	if ctx.PrimitiveType() != nil {
		symbolType = t.at.GetTypeOfNode()[ctx.PrimitiveType()]
		t.at.PutTypeOfNode(ctx, symbolType)
	} else if ctx.ClassOrInterfaceType() != nil {
		symbolType = t.at.GetTypeOfNode()[ctx.ClassOrInterfaceType()]
		t.at.PutTypeOfNode(ctx, symbolType)
	} else if ctx.FunctionType() != nil {
		symbolType = t.at.GetTypeOfNode()[ctx.FunctionType()]
		t.at.PutTypeOfNode(ctx, symbolType)
	}
	if symbolType == nil {
		return
	}
	// 还原为默认值
	symbolType.SetArray(false)
	if len(ctx.AllLBRACK()) > 0 && len(ctx.AllRBRACK()) > 0 {
		symbolType.SetArray(true)

	}
}

func (t *TypeResolver) EnterClassOrInterfaceType(ctx *parser.ClassOrInterfaceTypeContext) {
	scope := t.at.FindEncloseScopeOfNode(ctx)
	class := t.at.FindClass(scope, ctx.GetText())
	t.at.PutTypeOfNode(ctx, class)
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
	if ctx.ANY() != nil {
		symbolType = symbol.Any
	}
	if ctx.BYTE() != nil {
		symbolType = symbol.Byte
	}
	t.at.PutTypeOfNode(ctx, symbolType)
}

// ExitFunctionType 函数类型
func (t *TypeResolver) ExitFunctionType(ctx *parser.FunctionTypeContext) {
	returnType := t.at.GetTypeOfNode()[ctx.TypeTypeOrVoid()]
	declareFunctionType := symbol.NewDeclareFunctionType(returnType)
	t.at.AppendType(declareFunctionType)
	t.at.PutTypeOfNode(ctx, declareFunctionType)

	if ctx.TypeList() != nil {
		for _, context := range ctx.TypeList().(*parser.TypeListContext).AllTypeType() {
			t := t.at.GetTypeOfNode()[context]
			declareFunctionType.AppendParameterType(t)
		}
	}
}
