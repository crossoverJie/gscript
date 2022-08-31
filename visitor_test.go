package gscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGScriptVisitor_Visit_Lexer(t *testing.T) {
	expression := "(2+3) * 2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q) %d\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText(), t.GetColumn())
	}
}
func TestGScriptVisitor_Visit(t *testing.T) {
	expression := "(2+3) * 2;"
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, "=", result)
	assert.Equal(t, result, 10)
}
func TestGScriptVisitor_Visit2(t *testing.T) {
	expression := "5%2;"
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, "=", result)
	assert.Equal(t, result, 1)
}
func TestMod(t *testing.T) {
	val, err := strconv.ParseInt("2", 32, 10)
	fmt.Println(val, err)
	fmt.Println(5 % 2)
}

func TestGScriptVisitor_VisitIfElse5(t *testing.T) {
	expression := `
if(3==(1+2)){
	return 1+2*3;
}`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, 7)
}
func TestGScriptVisitor_VisitIfElse6(t *testing.T) {
	expression := `
if(3<(1+2)){
	return 1+2*3;
} else {
	return 2;
}`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, int(2))
}
func TestGScriptVisitor_VisitIfElse7(t *testing.T) {
	expression := `
if(3<(1+2)){
	return 1+2*3;
} else {
	return "123";
}`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, "123")
}
func TestGScriptVisitor_VisitIfElse8(t *testing.T) {
	expression := `
if(3!=(1+2)){
	return 1+3;
} else {
	return false;
}`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, false)
}

func TestArithmeticOperators(t *testing.T) {
	expression := `
if(4==(2+2)){
	return 1+3;
} else {
	return false;
}`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, int(4))
}
func TestArithmeticOperators2(t *testing.T) {
	expression := `(10+20)*20;`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, 600)
}
func TestArithmeticOperators3(t *testing.T) {
	expression := `(1+1.1)-2;`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result.(float64))

	expression = `(10+10)*10+10.1;`
	result = NewCompiler().Compiler(expression)
	fmt.Println(result)
}
func TestArithmeticOperators4(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return true ;
} else {
	return 20 ;
}
`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, true)

}
func TestArithmeticOperators5(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return 10++ ;
} else {
	return 20 ;
}
`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, 11)
}
func TestArithmeticOperators6(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return !(1+1==2) ;
} else {
	return 20 ;
}
`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, false)
	expression = `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) ;
} else {
	return 20 ;
}
`
	result = NewCompiler().Compiler(expression)
	fmt.Println(result)
	assert.Equal(t, result, true)
}
func TestArithmeticOperators7(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return !10 ;
} else {
	return 20 ;
}
`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
}

func TestDeclare(t *testing.T) {
	expression := `int a=100;`
	var result = NewCompiler().Compiler(expression)
	fmt.Println(result)
}
