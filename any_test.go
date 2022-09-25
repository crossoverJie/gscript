package gscript

import (
	"testing"
)

func TestAny(t *testing.T) {
	script := `
any a =10;
println(a);

int fun1(any a,int b){
	return a+b;
}
int v =fun1(1,2);
println(v);
assertEqual(v,3);

any v2 = fun1(1,2);
println(v2);
assertEqual(v2,3);

int fun2(int a, any b){
	return a+b;
}
int v3 =fun2(1,2);
println(v3);
assertEqual(v3,3);

int fun3(any a, any b){
	return a+b;
}
int v4 =fun3(1,2);
println(v4);
assertEqual(v4,3);

int fun4(any a, any b){
	return a+b;
}
string v5 =fun4("10", "20");
println(v5);
assertEqual(v5,"1020");
`
	NewCompiler().Compiler(script)
}
func TestAny2(t *testing.T) {
	script := `
	any[] a = {1,2,3};
	println(a);
	println(len(a));
	assertEqual(3, len(a));
	
	int[] b = {2,3,4};
	println(b);
	println(len(b));
	assertEqual(3, len(b));
`
	NewCompiler().Compiler(script)
}
func TestAny4(t *testing.T) {
	script := `
any a = "1";
any b = "1";
println(a==b);
assertEqual(a==b,true);

int x1 = 1;
int x2 =x1;
println(x1==x2);
assertEqual(x1==x2,true);

any c = a;
println(a==c);
assertEqual(a==c,true);
println(a!=c);
assertEqual(a!=c,false);
`
	NewCompiler().Compiler(script)
}
