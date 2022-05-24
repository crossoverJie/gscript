package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitStatus(t *testing.T) {
	str := "age >= 45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Identifier)
	assert.Equal(t, types[0].value, "age")
	assert.Equal(t, types[1].tokenType, GE)
	assert.Equal(t, types[1].value, ">=")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "45")
}
func TestInitStatus2(t *testing.T) {
	str := "age>=45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Identifier)
	assert.Equal(t, types[0].value, "age")
	assert.Equal(t, types[1].tokenType, GE)
	assert.Equal(t, types[1].value, ">=")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "45")
}

func TestInitStatus3(t *testing.T) {
	str := "age>45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Identifier)
	assert.Equal(t, types[0].value, "age")
	assert.Equal(t, types[1].tokenType, GT)
	assert.Equal(t, types[1].value, ">")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "45")
}
func TestInitStatus4(t *testing.T) {
	str := "age > 45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Identifier)
	assert.Equal(t, types[0].value, "age")
	assert.Equal(t, types[1].tokenType, GT)
	assert.Equal(t, types[1].value, ">")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "45")
}
func TestInitStatus50(t *testing.T) {
	str := "int a = b+3"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Int)
	assert.Equal(t, types[0].value, "int")
	assert.Equal(t, types[1].tokenType, Identifier)
	assert.Equal(t, types[1].value, "a")
	assert.Equal(t, types[2].tokenType, Assignment)
	assert.Equal(t, types[2].value, "=")
	assert.Equal(t, types[3].tokenType, Identifier)
	assert.Equal(t, types[3].value, "b")
	assert.Equal(t, types[4].tokenType, Plus)
	assert.Equal(t, types[4].value, "+")
	assert.Equal(t, types[5].tokenType, IntLiteral)
	assert.Equal(t, types[5].value, "3")
}
func TestInitStatus51(t *testing.T) {
	str := "int a = (b+3)*2"
	Tokenize(str)
}
func TestInitStatus5(t *testing.T) {
	str := "int a = 45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Int)
	assert.Equal(t, types[0].value, "int")
	assert.Equal(t, types[1].tokenType, Identifier)
	assert.Equal(t, types[1].value, "a")
	assert.Equal(t, types[2].tokenType, Assignment)
	assert.Equal(t, types[2].value, "=")
	assert.Equal(t, types[3].tokenType, IntLiteral)
	assert.Equal(t, types[3].value, "45")
}
func TestInitStatus6(t *testing.T) {
	str := "inta = 45"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, Identifier)
	assert.Equal(t, types[0].value, "inta")
	assert.Equal(t, types[1].tokenType, Assignment)
	assert.Equal(t, types[1].value, "=")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "45")
}
func TestInitStatus7(t *testing.T) {
	str := "2+3*5"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, IntLiteral)
	assert.Equal(t, types[0].value, "2")
	assert.Equal(t, types[1].tokenType, Plus)
	assert.Equal(t, types[1].value, "+")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "3")
	assert.Equal(t, types[3].tokenType, Star)
	assert.Equal(t, types[3].value, "*")
	assert.Equal(t, types[4].tokenType, IntLiteral)
	assert.Equal(t, types[4].value, "5")
}
func TestInitStatus8(t *testing.T) {
	str := "2+3*5+7+9*10"
	types := Tokenize(str)
	assert.Equal(t, types[0].tokenType, IntLiteral)
	assert.Equal(t, types[0].value, "2")
	assert.Equal(t, types[1].tokenType, Plus)
	assert.Equal(t, types[1].value, "+")
	assert.Equal(t, types[2].tokenType, IntLiteral)
	assert.Equal(t, types[2].value, "3")
	assert.Equal(t, types[3].tokenType, Star)
	assert.Equal(t, types[3].value, "*")
	assert.Equal(t, types[4].tokenType, IntLiteral)
	assert.Equal(t, types[4].value, "5")
}
func TestInitStatus9(t *testing.T) {
	str := "2+3*5+7+9*10\n int b = 10\n"
	Tokenize(str)
}
