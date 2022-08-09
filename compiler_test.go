package gscript

import "testing"

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
