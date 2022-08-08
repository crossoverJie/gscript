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
	assert.Equal(t, peek.String(), "sex")

	pop := s.Pop()
	assert.Equal(t, pop.String(), "sex")
	pop = s.Pop()
	assert.Equal(t, pop.String(), "age")
	pop = s.Pop()
	assert.Equal(t, pop.String(), "name")

	empty := s.IsEmpty()
	assert.Equal(t, empty, true)
}
