package symbol

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/container"
)

type Symbol interface {
	SetName(name string)
	GetName() string
	GetEncloseScope() Scope
	SetEncloseScope(scope Scope)
}

type symbol struct {
	name string

	// 所属的作用域
	encloseScope Scope

	// 可见性
	visibility int
	ctx        antlr.ParserRuleContext
}

func (s *symbol) SetName(name string) {
	s.name = name
}
func (s *symbol) GetName() string {
	return s.name
}
func (s *symbol) GetEncloseScope() Scope {
	return s.encloseScope
}

func (s *symbol) SetEncloseScope(scope Scope) {
	s.encloseScope = scope
}

type Type interface {
	GetName() string
	GetEncloseScope() Scope

	// IsType 判断类型是否相同，或者是不是 is-a ，是否为子类
	IsType(t Type) bool
}

// MatchNil nil 可以满足所有类型，除了基本类型
func MatchNil(t Type) bool {
	primitiveType, ok := t.(*PrimitiveType)
	// nil 匹配所有类型
	if ok && primitiveType.GetName() == "nil" {
		return true
	}
	return false
}

type Scope interface {
	Symbol
	AddSymbol(symbol Symbol)
	ContainsSymbol(symbol Symbol) bool
	GetSymbols() []Symbol
	SetCtx(ctx antlr.ParserRuleContext)
	GetCtx() antlr.ParserRuleContext
	GetVariable(name string) *Variable
	GetFunction(name string, paramTypes []Type) *Func
	GetFunctionVariable(name string, paramTypes []Type) *Variable
	GetClass(name string) *Class
	String() string
}

type scope struct {
	*symbol
	symbols []Symbol
}

func (s *scope) AddSymbol(symbol Symbol) {
	s.symbols = append(s.symbols, symbol)
	symbol.SetEncloseScope(s)
}
func (s *scope) ContainsSymbol(symbol Symbol) bool {
	for _, s := range s.symbols {
		if s == symbol {
			return true
		}
	}
	return false
}
func (s *scope) GetSymbols() []Symbol {
	return s.symbols
}
func (s *scope) SetCtx(ctx antlr.ParserRuleContext) {
	s.ctx = ctx
}
func (s *scope) GetCtx() antlr.ParserRuleContext {
	return s.ctx
}
func (s *scope) GetVariable(name string) *Variable {
	return GetVariable(s, name)
}

func (s *scope) GetFunction(name string, paramTypes []Type) *Func {
	return GetFunction(s, name, paramTypes)
}

func (s *scope) GetFunctionVariable(name string, paramTypes []Type) *Variable {
	return GetFunctionVariable(s, name, paramTypes)
}

func (s *scope) GetClass(name string) *Class {
	return GetClass(s, name)
}

// GetVariable 从 scope 中通过变量名称查询变量
func GetVariable(scope Scope, name string) *Variable {
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *Variable:
			if s.GetName() == name {
				return s.(*Variable)
			}
		}
	}
	return nil
}

// GetFunctionVariable 在当前 scope 中查找函数变量，其中变量中的函数的入参得匹配。
func GetFunctionVariable(scope Scope, name string, paramTypes []Type) *Variable {
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *Variable:
			variable := s.(*Variable)
			funcType, ok := variable.GetType().(FuncType)
			if variable.GetName() == name && ok && funcType.MatchParameterTypes(paramTypes) {
				return variable
			}
		}
	}
	return nil
}

// GetFunction 在 scope 中查询函数，需要入参相同
func GetFunction(scope Scope, name string, paramTypes []Type) *Func {
	// todo crossoverJie 返回值校验
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *Func:
			f := s.(*Func)
			if name == f.GetName() && f.MatchParameterTypes(paramTypes) {
				return f
			}
		}
	}

	return nil
}

// GetClass 在 scope 中查找 class 的 symbol
func GetClass(scope Scope, name string) *Class {
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *Class:
			c := s.(*Class)
			if c.GetName() == name {
				return c
			}
		}
	}
	return nil
}

func (s *scope) String() string {
	return "scope:" + s.GetName()
}

// todo crossoverJie thread-safe?
var blockIndex int

type BlockScope struct {
	*scope
}

func NewBlockScope(ctx antlr.ParserRuleContext, name string, s Scope) Scope {
	blockIndex++
	name = fmt.Sprintf("block-%s%d", name, blockIndex)
	blockScope := &BlockScope{
		scope: &scope{
			symbols: make([]Symbol, 0),
			symbol: &symbol{
				name:         name,
				encloseScope: s,
				visibility:   0,
				ctx:          ctx,
			},
		},
	}

	// todo crossoverJie 这段是否可以去掉？
	blockScope.SetCtx(ctx)
	blockScope.SetEncloseScope(s)
	blockScope.SetName(name)
	return blockScope
}

type Variable struct {
	*symbol
	t              Type
	isVariableArgs bool
}

func NewVariable(ctx antlr.ParserRuleContext, name string, enclose Scope, isVariableArgs bool) *Variable {
	return &Variable{
		symbol: &symbol{
			name:         name,
			encloseScope: enclose,
			ctx:          ctx,
		},
		isVariableArgs: isVariableArgs,
	}
}

func (v *Variable) GetType() Type {
	return v.t
}
func (v *Variable) SetType(t Type) {
	v.t = t
}
func (v *Variable) String() string {
	return v.name
}
func (v *Variable) getVariableArgs() bool {
	return v.isVariableArgs
}

// function

type FuncType interface {
	Type
	GetReturnType() Type

	GetParameterType() []Type

	// MatchParameterTypes 检查参数类型是否和函数匹配
	MatchParameterTypes(types []Type) bool
}

type Func struct {
	*scope
	parameters    []*Variable
	parameterType []Type

	// todo crossoverJie 返回多个值
	returnType Type

	//闭包变量，用于存放外部变量
	closureVar container.Set
}

func NewFunc(ctx antlr.ParserRuleContext, name string, encloseScope Scope) *Func {
	return &Func{
		scope: &scope{
			symbol: &symbol{
				name:         name,
				encloseScope: encloseScope,
				ctx:          ctx,
			},
		},
	}
}
func (f *Func) SetClosureVar(closureVar container.Set) {
	f.closureVar = closureVar
}
func (f *Func) GetClosureVar() container.Set {
	return f.closureVar
}
func (f *Func) GetEncloseScope() Scope {
	return f.encloseScope
}

func (f *Func) IsType(t Type) bool {
	if MatchNil(t) {
		return true
	}
	funcType, ok := t.(FuncType)
	if ok {
		return isFuncType(f, funcType)
	}
	return f == t
}

func (f *Func) GetReturnType() Type {
	return f.returnType
}

func (f *Func) SetReturnType(t Type) {
	f.returnType = t
}

// GetParameterType 获取和初始化
func (f *Func) GetParameterType() []Type {
	if f.parameterType == nil {
		f.parameterType = make([]Type, 0)
	}
	for _, parameter := range f.parameters {
		f.parameterType = append(f.parameterType, parameter.GetType())
	}

	return f.parameterType
}

func (f *Func) AppendParameter(v *Variable) {
	f.parameters = append(f.parameters, v)
}

// MatchParameterTypes 通过参数类型匹配函数是否一致
func (f *Func) MatchParameterTypes(paramTypes []Type) bool {
	if len(f.parameters) > 0 {
		last := f.parameters[len(f.parameters)-1]
		if last.isVariableArgs {
			// 存在可变参数
			// 比较可变参数前的普通参数
			normalParams := f.parameters[0 : len(f.parameters)-1]
			for i, normalParam := range normalParams {
				t := paramTypes[i]
				if !t.IsType(normalParam.t) {
					return false
				}
			}
			// 比较剩下类型与可变参数是否匹配
			types := paramTypes[len(normalParams):]
			for _, t := range types {
				if !t.IsType(last.t) {
					return false
				}
			}

			// 全部比较完成
			return true
		}
	}

	// 比较没有可变参数
	if len(f.parameters) != len(paramTypes) {
		return false
	}
	match := true
	for i, t := range paramTypes {
		variable := f.parameters[i]
		if !variable.t.IsType(t) {
			match = false
			break
		}
	}
	return match

}

func (f *Func) SetName(name string) {
	f.name = name
}
func (f *Func) GetName() string {
	return f.name
}

func (f *Func) SetEncloseScope(scope Scope) {
	f.encloseScope = scope
}

// IsConstructor 当前函数是否为构造函数，也就是函数名称是否与函数声明所在的 class 名称相同
func (f *Func) IsConstructor() bool {
	class, ok := f.scope.encloseScope.(*Class)
	if ok && class.GetName() == f.name {
		return true
	}
	return false
}

// IsMethod 是否是类的函数
func (f *Func) IsMethod() bool {
	_, ok := f.encloseScope.(*Class)
	return ok
}

var functionTypeIndex int

type DeclareFunctionType struct {
	name          string
	encloseScope  Scope
	returnType    Type   // 返回值类型
	parameterType []Type // 参数类型
}

// NewDeclareFunctionType 函数类型的变量
// func int(int) f2 = f1(); f2 的类型
func NewDeclareFunctionType(returnType Type) *DeclareFunctionType {
	functionTypeIndex = functionTypeIndex + 1
	name := fmt.Sprintf("DeclareFunctionType%d", blockIndex)
	return &DeclareFunctionType{name: name, returnType: returnType}
}
func (d *DeclareFunctionType) GetName() string {
	return d.name
}

func (d *DeclareFunctionType) GetEncloseScope() Scope {
	return d.encloseScope
}

func (d *DeclareFunctionType) IsType(t Type) bool {
	if MatchNil(t) {
		return true
	}
	funcType, ok := t.(FuncType)
	if ok {
		return isFuncType(d, funcType)
	}
	return false
}

func (d *DeclareFunctionType) GetReturnType() Type {
	return d.returnType
}

func (d *DeclareFunctionType) GetParameterType() []Type {
	return d.parameterType
}
func (d *DeclareFunctionType) AppendParameterType(t Type) {
	d.parameterType = append(d.parameterType, t)
}

func (d *DeclareFunctionType) MatchParameterTypes(types []Type) bool {
	if len(types) != len(d.parameterType) {
		return false
	}
	for i, t := range d.parameterType {
		if !t.IsType(types[i]) {
			return false
		}
	}

	return true
}

// 两个 FuncType 是否相等
func isFuncType(t1, t2 FuncType) bool {
	if t2 == t1 {
		return true
	}

	// 兼容直接传递函数变量的情况
	// 传递的是 func void(int, int) x = handle1，不需要 x
	// handleFunc("/abc", handle1);
	df, ok := t2.GetReturnType().(*DeclareFunctionType)
	if ok {
		returnType := isFuncType(t1, df)
		if !returnType {
			return false
		}
	} else {
		// 传递的是 func void(int, int) x = handle1
		// handleFunc("/abc", x);
		if t1.GetReturnType() != nil && t2.GetReturnType() != nil {
			// 函数变量的返回值不写时，也没写 void
			if !t1.GetReturnType().IsType(t2.GetReturnType()) {
				return false
			}
		}
	}

	if len(t1.GetParameterType()) != len(t2.GetParameterType()) {
		return false
	}
	for i, t := range t1.GetParameterType() {
		if !t.IsType(t2.GetParameterType()[i]) {
			return false
		}
	}

	return true
}
