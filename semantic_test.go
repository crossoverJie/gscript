package gscript

import "testing"

func TestReturn(t *testing.T) {
	script := `
int getNumber(){
	return 1;
}
`
	NewCompiler().Compiler(script)
}
