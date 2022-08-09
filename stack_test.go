package gscript

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Peek(t *testing.T) {
	s := &Stack{}
	s.Push(NewStackFrame("name"))
	s.Push(NewStackFrame("age"))
	s.Push(NewStackFrame("sex"))

	peek := s.Peek()
	assert.Equal(t, peek.(*stackFrame).String(), "sex")

	pop := s.Pop()
	assert.Equal(t, pop.(*stackFrame).String(), "sex")
	pop = s.Pop()
	assert.Equal(t, pop.(*stackFrame).String(), "age")
	pop = s.Pop()
	assert.Equal(t, pop.(*stackFrame).String(), "name")

	empty := s.IsEmpty()
	assert.Equal(t, empty, true)
}
