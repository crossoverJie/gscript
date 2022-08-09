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
	expression := "(2+3) * 2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	tree := parser.Parse()

	visitor := GScriptVisitor{}

	var result = visitor.Visit(tree)
	fmt.Println(expression, "=", result)
	assert.Equal(t, result, 10)
}
func TestGScriptVisitor_Visit2(t *testing.T) {
	expression := "5%2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()

	visitor := GScriptVisitor{}

	var result = visitor.Visit(tree)
	fmt.Println(expression, "=", result)
	assert.Equal(t, result, 1)
}
func TestMod(t *testing.T) {
	val, err := strconv.ParseInt("2", 32, 10)
	fmt.Println(val, err)
	fmt.Println(5 % 2)
}

func TestGScriptVisitor_VisitIfElse(t *testing.T) {
	expression := "1<=2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()

	visitor := GScriptVisitor{}

	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, true)
}
func TestGScriptVisitor_VisitIfElse2(t *testing.T) {
	expression := "1>=2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, false)
}
func TestGScriptVisitor_VisitIfElse3(t *testing.T) {
	expression := "1==2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, false)
}
func TestGScriptVisitor_VisitIfElse4(t *testing.T) {
	expression := "2==2"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, true)
}
func TestGScriptVisitor_VisitIfElse5(t *testing.T) {
	expression := `
if(3==(1+2)){
	return 1+2*3
}`
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Prog()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, 7)
}
func TestGScriptVisitor_VisitIfElse6(t *testing.T) {
	expression := `
if(3<(1+2)){
	return 1+2*3
} else {
	return 2
}`
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Prog()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, int(2))
}
func TestGScriptVisitor_VisitIfElse7(t *testing.T) {
	expression := `
if(3<(1+2)){
	return 1+2*3
} else {
	return "123"
}`
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Prog()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, "123")
}
func TestGScriptVisitor_VisitIfElse8(t *testing.T) {
	expression := `
if(3!=(1+2)){
	return 1+3
} else {
	return false
}`
	input := antlr.NewInputStream(expression)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGScriptParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Prog()
	visitor := GScriptVisitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, " result:", result)
	assert.Equal(t, result, false)
}

func TestArithmeticOperators(t *testing.T) {
	expression := `
if(4==(2+2)){
	return 1+3
} else {
	return false
}`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, int(4))
}
func TestArithmeticOperators2(t *testing.T) {
	expression := `(10+20)*20`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, 600)
}
func TestArithmeticOperators3(t *testing.T) {
	expression := `(1+1.1)-2`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret.(float64))

	expression = `(10+10)*10+10.1`
	ret = ArithmeticOperators(expression)
	fmt.Println(ret)
}
func TestArithmeticOperators4(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return true 
} else {
	return 20 
}
`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, true)

}
func TestArithmeticOperators5(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return 10++ 
} else {
	return 20 
}
`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, 11)
}
func TestArithmeticOperators6(t *testing.T) {
	expression := `
if ( (10 +10 ) == 20 ) {
	return !(1+1==2) 
} else {
	return 20 
}
`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, false)
	expression = `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) 
} else {
	return 20 
}
`
	ret = ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, true)
}
func TestArithmeticOperators7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
		}
	}()
	expression := `
if ( (10 +10 ) == 20 ) {
	return !10 
} else {
	return 20 
}
`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
	assert.Equal(t, ret, false)
}

func TestDeclare(t *testing.T) {
	expression := `int a=100`
	ret := ArithmeticOperators(expression)
	fmt.Println(ret)
}
