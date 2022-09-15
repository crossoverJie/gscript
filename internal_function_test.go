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
	args := []interface{}{
		"1", "2",
	}
	fmt.Printf("hello %s %s", args...)
	fmt.Println(fmt.Sprintf("abc %d", 123))
}
