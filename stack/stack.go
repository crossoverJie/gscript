package stack

import (
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

func NewObjectStackFrame(object Object) *Frame {
	return &Frame{object: object}
}
func NewBlockScopeFrame(scope symbol.Scope) *Frame {
	return &Frame{scope: scope, object: NewEmptyObject()}
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

func (s *Frame) String() string {
	return fmt.Sprintf("%s", s.object)
}

type Object interface {
	GetValue(variable *symbol.Variable) interface{}
	SetValue(variable *symbol.Variable, value interface{})
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
