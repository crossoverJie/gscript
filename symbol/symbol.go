package symbol

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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

type Scope interface {
	Symbol
	AddSymbol(symbol Symbol)
	ContainsSymbol(symbol Symbol) bool
	GetSymbols() []Symbol
	SetCtx(ctx antlr.ParserRuleContext)
	GetVariable(name string) *Variable
	GetFunction(name string, paramTypes []Type) *Func
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

func GetFunction(scope Scope, name string, paramTypes []Type) *Func {
	for _, s := range scope.GetSymbols() {
		switch s.(type) {
		case *Func:
			f := s.(*Func)
			if f.MatchParameterTypes(paramTypes) && name == f.GetName() {
				return f
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
	t Type
}

func NewVariable(ctx antlr.ParserRuleContext, name string, enclose Scope) *Variable {
	return &Variable{
		symbol: &symbol{
			name:         name,
			encloseScope: enclose,
			ctx:          ctx,
		},
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

	// todo crossoverJie 闭包变量，用于存放外部变量
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

func (f *Func) GetEncloseScope() Scope {
	return f.encloseScope
}

func (f *Func) IsType(t Type) bool {
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
