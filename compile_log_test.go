package gscript

import (
	"fmt"
	"github.com/crossoverJie/gscript/symbol"
	"testing"
)

func TestCompileFail(t *testing.T) {
	script := `
a+2;
b+c;
`
	NewCompiler().Compiler(script)

}
func TestCompileFail2(t *testing.T) {
	script := `
class T{
}
T t = T();
t.print();
`
	NewCompiler().Compiler(script)

}
func TestCompileFail3(t *testing.T) {
	script := `
class T{
	T(int a){
		
	}
}
T t = T("abc",123);
`
	NewCompiler().Compiler(script)
}
func TestCompileFail4(t *testing.T) {
	script := `
func1(1);
`
	NewCompiler().Compiler(script)
}
func TestCompileFail5(t *testing.T) {
	script := `
class T{
}
T t= T();
t.a;
`
	NewCompiler().Compiler(script)
}
func TestCompileFail6(t *testing.T) {
	script := `
println("1"-"2");
println("1"%"2");
`
	NewCompiler().Compiler(script)
}
func TestCompileFail7(t *testing.T) {
	script := `
int run(){}
int run(){}
`
	NewCompiler().Compiler(script)
}
func TestCompileFail8(t *testing.T) {
	script := `
class T{}
class T{}
`
	NewCompiler().Compiler(script)
}

func TestSlice2(t *testing.T) {
	types := get()
	for _, t := range types {
		if t == nil {
			continue
		}
		fmt.Println(t.GetName())
	}
}
func get() []symbol.Type {
	var paraTypes []symbol.Type
	paraTypes = append(paraTypes, nil)
	return paraTypes
}
