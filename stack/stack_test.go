package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Peek(t *testing.T) {
	s := &Stack{}
	s.Push("name")
	s.Push("age")
	s.Push("sex")
	assert.Equal(t, s.Size(), 3)

	peek := s.Peek()
	assert.Equal(t, peek, "sex")
	assert.Equal(t, s.Size(), 3)

	pop := s.Pop()
	assert.Equal(t, pop, "sex")
	assert.Equal(t, s.Size(), 2)
	pop = s.Pop()
	assert.Equal(t, pop, "age")
	assert.Equal(t, s.Size(), 1)

	pop = s.Pop()
	assert.Equal(t, pop, "name")
	assert.Equal(t, s.Size(), 0)

	empty := s.IsEmpty()
	assert.Equal(t, empty, true)
}

func TestStack_Get(t *testing.T) {
	s := &Stack{}
	s.Push("1")
	s.Push("2")
	s.Push("3")
	for i := s.Size() - 1; i >= 0; i-- {
		fmt.Println(s.Get(i))
	}
	fmt.Println()
	s.Pop()
	for i := s.Size() - 1; i >= 0; i-- {
		fmt.Println(s.Get(i))
	}

}
