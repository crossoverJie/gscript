package gscript

import (
	"fmt"
	"testing"
)

func TestVariableArgs(t *testing.T) {
	script := `
int get(int ...a){
	println(a);
	return a;
}
int[] b = get(10);
println(b);
int x = b[0];
assertEqual(x,10);
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestVariableArgs2(t *testing.T) {
	script := `
int get(int ...a){
	println(a);
	return a;
}
int[] b = get(10,2);
println(b);


string get2(string a, int ...b){
	println(a);
	println(b);
	int x = b[0];
	assertEqual(x,2);
	for(int i=0;i<len(b);i++){
		int x = b[i];
		assertEqual(x,i+2);
		println(x);
	}
	return a;
}
string g = get2("1",2,3,4,5);
println(g);
assertEqual(g, "1");
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestVariableArgs3(t *testing.T) {
	script := `

string get2(string a,float c, int ...b){
	println(a);
	println(b);
	println(c);
	int x = b[0];
	assertEqual(x,2);
	for(int i=0;i<len(b);i++){
		int x = b[i];
		assertEqual(x,i+2);
		println(x);
	}
	return a;
}
string g = get2("1",999.99, 2,3,4,5);
//println(g);
assertEqual(g, "1");
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestVariableArgs4(t *testing.T) {
	script := `

int add(string s, int ...num){
	println(s);
	int sum = 0;
	for(int i=0;i<len(num);i++){
		int v = num[i];
		sum = sum+v;
	}
	return sum;
}
int x = add("abc", 1,2,3,4);
println(x);
assertEqual(x, 10);
`
	NewCompiler().CompilerWithoutNative(script)
}

func TestVariable(t *testing.T) {
	test("123", 4, 5, 6)
}
func test(a string, b ...int) {
	fmt.Println(a, b)
	x := b[0 : len(b)-1]
	fmt.Println(x)

	ints := b[1:]
	fmt.Println(ints)
}
