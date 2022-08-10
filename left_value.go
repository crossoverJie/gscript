package gscript

import (
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
)

type LeftValue struct {
	variable *symbol.Variable
	object   stack.Object
}

func NewLeftValue(variable *symbol.Variable, object stack.Object) *LeftValue {
	return &LeftValue{
		variable: variable,
		object:   object,
	}
}

func (l *LeftValue) GetValue() interface{} {
	// todo crossoverJie this, super.

	return l.object.GetValue(l.variable)
}

func (l *LeftValue) SetValue(value interface{}) {
	l.object.SetValue(l.variable, value)

	// todo crossoverJie variable 是函数类型
}

func (l *LeftValue) String() string {
	return l.object.String()
}
