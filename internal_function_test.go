package gscript

import (
	"fmt"
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

func TestPrint(t *testing.T) {
	fmt.Printf("hello %s", "gscript")
	fmt.Println(fmt.Sprintf("abc %d", 123))
}
