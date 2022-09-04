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

	// variable 是函数类型
	funcObject, ok := value.(*stack.FuncObject)
	if ok {
		funcObject.SetReferenceVariable(l.variable)
	}
}
func (l *LeftValue) AppendArray(v interface{}) {
	array, ok := l.GetValue().([]interface{})
	if ok {
		array = append(array, v)
	} else {
		// todo crossoverJie 运行时报错
	}
}

func (l *LeftValue) String() string {
	return l.object.String()
}
