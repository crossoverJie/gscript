package gscript

import (
	"os"
	"testing"
)

func TestByte(t *testing.T) {
	script := `
string v5="9898";
`
	//os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestByte2(t *testing.T) {
	script := `
int[] a={1,2};
println(a);
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
