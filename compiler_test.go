package gscript

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompiler_Compiler(t *testing.T) {
	script := `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) 
} else {
	return 20 
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Compiler2(t *testing.T) {
	script := `
int a=10
a++
return a
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 11)
}
func TestCompiler_CompilerFor(t *testing.T) {
	script := `
int age = 0 
for(int i = 0;i<100;i++) {
	age++
} 
return age
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 100)
}
func TestCompiler_CompilerFor3(t *testing.T) {
	script := `
int age = 0 
int sum=1
for(int i = 0;i<100;i++) {
	sum=sum+1
	for(int i = 0;i<100;i++) {
		age++
	}
} 
return age
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 100*100)
}
func TestCompiler_CompilerFor2(t *testing.T) {
	script := `
int age = 1 
for(int i = 0;i<100;i++) {
	age=age+1
} 
return age
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 101)
}
func TestCompiler_CompilerIf(t *testing.T) {
	script := `
int age=10
if (age>10){
	age++
}else{
	age--
}
return age
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 9)
}
func TestCompiler_CompilerIf2(t *testing.T) {
	script := `
int age=10
if (age>10){
	age++
}else{
	age--
}
return abc
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}
