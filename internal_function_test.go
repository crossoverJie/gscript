package gscript

import (
	"testing"
)

func TestFunction(t *testing.T) {
	script := `
int a(){
	return 10;
}
println(a());
if (1<a()){
	println("1<a()");
}
`
	NewCompiler().Compiler(script)
}
