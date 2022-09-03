package stack

import (
	"bytes"
	"fmt"
	"github.com/crossoverJie/gscript/symbol"
)

type Stack []interface{}

// IsEmpty  check if Stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
func (s *Stack) Get(i int) interface{} {
	return (*s)[i]
}

// Push a new value onto the Stack
func (s *Stack) Push(val interface{}) {
	*s = append(*s, val) // Simply append the new value to the end of the Stack
}

// Pop Remove and return top element of Stack. Return false if Stack is empty.
func (s *Stack) Pop() interface{} {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the Stack by slicing it off.
	return element

}

// Peek Peek value, not remove element.
func (s *Stack) Peek() interface{} {
	index := len(*s) - 1
	element := (*s)[index] // Index into the slice and obtain the element.
	return element
}

type Frame struct {
	scope  symbol.Scope
	parent *Frame
	object Object
}

func NewObjectStackFrame(object *FuncObject) *Frame {
	return &Frame{object: object, scope: object.GetFunction()}
}
func NewBlockScopeFrame(scope symbol.Scope) *Frame {
	return &Frame{scope: scope, object: NewEmptyObject()}
}
func NewClassStackFrame(classObject *ClassObject) *Frame {
	return &Frame{scope: classObject.class, object: classObject}
}

func (s *Frame) GetScope() symbol.Scope {
	return s.scope
}
func (s *Frame) GetObject() Object {
	return s.object
}

func (s *Frame) GetParent() *Frame {
	return s.parent
}
func (s *Frame) SetParent(f *Frame) {
	s.parent = f
}

func (s *Frame) String() string {
	return fmt.Sprintf("%s", s.object)
}

// ContainsVariable 当前栈帧中是否包含某个变量的数据
func (s *Frame) ContainsVariable(variable *symbol.Variable) bool {
	if s.object != nil {
		value := s.object.GetValue(variable)
		return value != nil
	}
	return false
}

// Object 变量对象，可以获取和保存变量的值
type Object interface {
	GetValue(variable *symbol.Variable) interface{}
	SetValue(variable *symbol.Variable, value interface{})
	String() string
}

type object struct {
	fields map[*symbol.Variable]interface{}
}

func NewEmptyObject() Object {
	return &object{fields: make(map[*symbol.Variable]interface{})}
}

func NewVariableObject() Object {
	return &object{fields: make(map[*symbol.Variable]interface{})}
}

func (o *object) GetValue(variable *symbol.Variable) interface{} {
	return o.fields[variable]
}
func (o *object) SetValue(variable *symbol.Variable, value interface{}) {
	o.fields[variable] = value
}
func (o *object) String() string {
	var b bytes.Buffer
	for k, v := range o.fields {
		b.WriteString(fmt.Sprintf("%s->%s \t", k.String(), v))
	}
	return fmt.Sprintf("%s", b.String())
}

type FuncObject struct {
	*object
	function *symbol.Func

	// 改函数被哪个变量所引用了，函数一等公民时。
	referenceVariable *symbol.Variable
}

func NewFuncObject(function *symbol.Func) *FuncObject {
	return &FuncObject{
		object:   &object{fields: make(map[*symbol.Variable]interface{})},
		function: function,
	}
}

func (f *FuncObject) GetFunction() *symbol.Func {
	return f.function
}

func (f *FuncObject) SetValue(variable *symbol.Variable, value interface{}) {
	f.fields[variable] = value
}

func (f *FuncObject) SetReferenceVariable(variable *symbol.Variable) {
	f.referenceVariable = variable
}

func (f *FuncObject) GetReferenceVariable() *symbol.Variable {
	return f.referenceVariable
}

type ClassObject struct {
	*object
	class *symbol.Class
}

func NewClassObject(class *symbol.Class) *ClassObject {
	return &ClassObject{
		object: &object{fields: make(map[*symbol.Variable]interface{})},
		class:  class,
	}
}

func (c *ClassObject) AllField() map[*symbol.Variable]interface{} {
	return c.fields
}

var (
	BreakObjectInstance    = &BreakObject{}
	ContinueObjectInstance = &ContinueObject{}
)

type BreakObject struct{}
type ContinueObject struct{}
