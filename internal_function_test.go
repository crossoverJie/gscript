package gscript

import (
	"testing"
)

func TestFunction(t *testing.T) {
	script := `
int a(){
	return 10;
}
print(a());
if (1<a()){
	print("1<a()");
}
`
	NewCompiler().Compiler(script)
}
