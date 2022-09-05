package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := NewSet()
	s.Add("1")
	s.Add("2")
	s.Add("3")
	add := s.Add("3")
	assert.Equal(t, add, false)
	assert.Equal(t, s.Size(), 3)

	s2 := NewSet()
	s2.Add("2")
	s2.Add("3")

	s.RemoveAll(s2)
	assert.Equal(t, s.Size(), 1)
	assert.Equal(t, s.List()[0], "1")

	s3 := NewSet()
	s3.Add("7")
	s3.Add("8")
	s.AddAll(s3)
	assert.Equal(t, s.Size(), 3)

}
