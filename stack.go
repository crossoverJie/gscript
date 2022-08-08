package gscript

import "fmt"

type Stack []*stackFrame

// IsEmpty  check if Stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the Stack
func (s *Stack) Push(val *stackFrame) {
	*s = append(*s, val) // Simply append the new value to the end of the Stack
}

// Pop Remove and return top element of Stack. Return false if Stack is empty.
func (s *Stack) Pop() *stackFrame {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the Stack by slicing it off.
	return element

}

// Peek Peek value, not remove element.
func (s *Stack) Peek() *stackFrame {
	index := len(*s) - 1
	element := (*s)[index] // Index into the slice and obtain the element.
	return element
}

type stackFrame struct {
	object interface{}
}

func NewStackFrame(obj interface{}) *stackFrame {
	return &stackFrame{object: obj}
}

func (s *stackFrame) String() string {
	return fmt.Sprintf("%s", s.object)
}
