package symbol

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Symbol struct {
	name string

	// 所属的作用域
	encloseScope *Scope

	// 可见性
	visibility int
	ctx        antlr.ParserRuleContext
}

func (s *Symbol) GetName() string {
	return s.name
}
func (s *Symbol) GetEncloseScope() *Scope {
	return s.encloseScope
}

type Type interface {
	GetName() string
	GetEncloseScope() *Scope
}

type Scope struct {
	symbols []*Symbol
	*Symbol
}

func (s *Scope) AddSymbol(symbol *Symbol) {
	s.symbols = append(s.symbols, symbol)
	symbol.encloseScope = s
}
