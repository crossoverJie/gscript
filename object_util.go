package gscript

import "fmt"

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

var (
	BreakObjectInstance    = &BreakObject{}
	ContinueObjectInstance = &ContinueObject{}
)

type BreakObject struct{}
type ContinueObject struct{}
type ReturnObject struct {
	returnObject interface{}
}

func NewReturnObject(object interface{}) *ReturnObject {
	return &ReturnObject{returnObject: object}
}

func (r *ReturnObject) GetReturnObject() interface{} {
	return r.returnObject
}
func (r *ReturnObject) String() string {
	return fmt.Sprintf("%d", r.returnObject)
}
