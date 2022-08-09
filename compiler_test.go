package gscript

import (
	"fmt"
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
int a = 10
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}
