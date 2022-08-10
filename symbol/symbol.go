package symbol

import "github.com/antlr/antlr4/runtime/Go/antlr"

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
	GetEncloseScope() *Scope
}

type Scope interface {
	Symbol
	AddSymbol(symbol Symbol)
	ContainsSymbol(symbol Symbol) bool
	GetSymbols() []Symbol
	SetCtx(ctx antlr.ParserRuleContext)
	GetVariable(name string) *Variable
	String() string
}

type scope struct {
	symbols []Symbol
	*symbol
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
func (s *scope) GetVariable(name string) *Variable {
	return GetVariable(s, name)
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

func (s *scope) String() string {
	return "scope:" + s.GetName()
}

type blockScope struct {
	*scope
	index int
}

func NewBlockScope(ctx antlr.ParserRuleContext, name string, s Scope) Scope {
	blockScope := &blockScope{
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
}

func NewVariable(ctx antlr.ParserRuleContext, name string, enclose Scope) *Variable {
	return &Variable{
		&symbol{
			name:         name,
			encloseScope: enclose,
			ctx:          ctx,
		},
	}
}

func (v *Variable) String() string {
	return v.name
}
