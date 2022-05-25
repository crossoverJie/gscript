package syntax

import (
	"fmt"
	"github.com/crossoverJie/gscript/token"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCalculator(t *testing.T) {
	script := "int a = b+3\n"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	declare, err := IntDeclare(reader)
	assert.Nil(t, err)
	log.Println("==============")
	declare.Print(Indent)
}
func TestCalculator2(t *testing.T) {
	script := "int a = 2+3*5\n"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	declare, err := IntDeclare(reader)
	assert.Nil(t, err)
	log.Println("==============")
	declare.Print(Indent)
}
func TestCalculator3(t *testing.T) {
	script := "int a = "
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	_, err := IntDeclare(reader)
	log.Println(err)
	assert.NotNil(t, err)
}
func TestCalculator4(t *testing.T) {
	script := "int a = 3+"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	_, err := IntDeclare(reader)
	log.Println(err)
	assert.NotNil(t, err)
}
func TestCalculator5(t *testing.T) {
	script := "int a = 3+(4+5"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	_, err := IntDeclare(reader)
	log.Println(err)
	assert.NotNil(t, err)
}
func TestCalculator6(t *testing.T) {
	script := "int a = 2+3*5+9\n"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	node, err := IntDeclare(reader)
	assert.Nil(t, err)
	node.Print(Indent)
}
func TestCalculator7(t *testing.T) {
	script := "int a = 45+100\n"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	node, err := IntDeclare(reader)
	assert.Nil(t, err)
	node.Print(Indent)
}

func TestCalculator8(t *testing.T) {
	script := "int a = 10+"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	_, err := IntDeclare(reader)
	fmt.Println(err)
	assert.NotNil(t, err)
}
func TestCalculator9(t *testing.T) {
	script := "int a = 10+\n"
	types := token.Tokenize(script)
	reader := NewTokenReader(types)
	_, err := IntDeclare(reader)
	fmt.Println(err)
	assert.NotNil(t, err)
}
