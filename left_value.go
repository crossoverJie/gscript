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

func (l *LeftValue) String() string {
	return l.object.String()
}

// todo crossoverJie 修改模块目录
type ArrayObject struct {
	left  *LeftValue
	index int
}

// NewArrayObject 数据对象
func NewArrayObject(left *LeftValue, index int) *ArrayObject {
	return &ArrayObject{
		left:  left,
		index: index,
	}
}
func (a *ArrayObject) GetLeftValue() *LeftValue {
	return a.left
}
func (a *ArrayObject) GetIndex() int {
	return a.index
}

// GetIndexValue 根据下标获取数组值
func (a *ArrayObject) GetIndexValue() interface{} {
	array := a.left.GetValue().([]interface{})
	return array[a.index]
}

// SetIndexValue 根据下标为数组赋值
func (a *ArrayObject) SetIndexValue(val interface{}) {
	array := a.left.GetValue().([]interface{})
	array[a.index] = val
}
