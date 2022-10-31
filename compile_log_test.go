package gscript

import (
	"fmt"
	"os"
	"testing"
)

func TestCompileFail(t *testing.T) {
	script := `
a+2;
b+c;
`
	NewCompiler().Compiler(script)

}
func TestByteFail(t *testing.T) {
	script := `
int[] a=10;
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestByteFail2(t *testing.T) {
	script := `
byte[] a=10;
println(a);
`
	os.Setenv(RuntimeError, "true")
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
int run(){return 0;}
int run(){return 0;}
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
func TestRuntimeFail1(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
println(10/0);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail2(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
println("1"/2);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail3(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
println("1"*2);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail4(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
string a=10;
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail5(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int a="s12";
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail6(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int a=10.1;
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail7(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int a=true;
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail8(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int a=nil;
println(a);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail9(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
class Person{
	int age;
	Person(int a){
		age = a;
	}
}
Person p1 = Person(10);
Person p2 = Person(20);
println(p1+p2);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail10(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int b=10;
for (b){
	println("1");
}
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail11(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
for (100){
	println("1");
}
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail12(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int a=10;
append(a,10);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail13(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
hash(12,1);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail14(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
int[] a ={1,2,3};
append(a,19,2);
`
	NewCompiler().Compiler(script)
}
func TestRuntimeFail15(t *testing.T) {
	os.Setenv(RuntimeError, "true")
	script := `
DateTime d = DateTime();
string local = d.getCurrentTime("Asia/Shanghai","2006-01-02 15:04:05");
printf("local:%s ", local);
int u = d.unix("Asia/Shanghai");
println(u);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}
