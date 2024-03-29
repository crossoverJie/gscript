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
	fmt.Printf("hello %s %s\n", args...)
	fmt.Println(fmt.Sprintf("abc %d", 123))
}

func TestPrintf(t *testing.T) {
	script := `

printf("hello %s ","123");
printf("hello-%s-%s ","123","abc");
printf("hello-%s-%d ","123",123);
string format = "this is %s ";
printf(format, "gscript");

string s = sprintf("nice to meet %s", "you");
println(s);
assertEqual(s,"nice to meet you");
`
	NewCompiler().Compiler(script)
}

func TestPrint2(t *testing.T) {
	script := `
print("123" + " ");
print("456" + " ");
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}
func TestDate2(t *testing.T) {
	script := `
DateTime d = DateTime();
string local = d.getCurrentTime("Asia/Shanghai","2006-01-02 15:04:05");
printf("local:%s", local);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}
