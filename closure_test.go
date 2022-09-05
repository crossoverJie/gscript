package gscript

import "testing"

func TestClosure(t *testing.T) {
	script := `
int varExternal =10;
func int(int) f1(){
	int varInner = 20;
	int innerFun(int a){
		println(a);
		int c=100;
		varExternal++;
		varInner++;
		return varInner;
	}
	return innerFun;
}

func int(int) f2 = f1();
for(int i=0;i<2;i++){
	println("varInner=" + f2(i) + ", varExternal=" + varExternal);
}
println("=======");
func int(int) f3 = f1();
for(int i=0;i<2;i++){
	println("varInner=" + f3(i) + ", varExternal=" + varExternal);
}

`
	NewCompiler().Compiler(script)
}
func TestClosure2(t *testing.T) {
	script := `
int varExternal =10;
func int() f1(){
	int varInner = 20;
	int innerFun(){
		return varInner;
	}
	return innerFun;
}

func int() f2 = f1();
println(f2());
assertEqual(f2(), 20);
`
	NewCompiler().Compiler(script)
}
func TestClosure3(t *testing.T) {
	script := `
string fun1(int a){
	println("fun1 a=" + a);
	return "test";
}

fun2 (func string(int) f){
    string ret = f(6);
    println("fun2 ret =" + ret);
}
func string(int) f = fun1;
fun2(f);
fun2(fun1);
`
	NewCompiler().Compiler(script)
}
func TestClosure5(t *testing.T) {
	script := `
class Test{
	fun2 (func string(int) f){
		string ret = f(6);
		println("fun2 ret =" + ret);
	}
}
string fun1(int a){
	println("fun1 a=" + a);
	return "test";
}


func string(int) f = fun1;
Test t = Test();
t.fun2(f);
`
	NewCompiler().Compiler(script)
}
func TestClosure6(t *testing.T) {
	script := `
class Test{
	func string(int) fun;
}
string fun1(int a){
	println("fun1 a=" + a);
	return "test";
}

fun2 (func string(int) f){
    string ret = f(6);
    println("fun2 ret =" + ret);
}


func string(int) f = fun1;
Test t = Test();
t.fun= fun1;
println(t.fun(100));
`
	NewCompiler().Compiler(script)
}
func TestClosure7(t *testing.T) {
	script := `
class Test{
	int value=0;
	Test(int v){
		value=v;
	}

	int map(func int(int) f){
		return f(value);
	}
}
int square(int v){
	return v*v; 
}
int add(int v){
	return v++; 
}
int add2(int v){
	v=v+2;
	return v; 
}
Test t =Test(100);
func int(int) s=square;
func int(int) a=add;
func int(int) a2=add2;
println(t.map(s));
assertEqual(t.map(s),10000);

println(t.map(a));
assertEqual(t.map(a),101);

println(t.map(a2));
assertEqual(t.map(a2),102);

// todo crossoverJie 这种情况编译报错
//println(t.map(add2));
//assertEqual(t.map(add2),102);
`
	NewCompiler().Compiler(script)
}

func TestClosure4(t *testing.T) {
	script := `
void fun(){
    int a = 0;
    int b = 1;
    int fibonacci(){
        int c = a;
        a = b;
        b = a+c;
        return c;
    }
    return fibonacci;
}

func int() fib = fun();

//打印斐波那契数列
for (int i = 0; i < 10; i++){
    println(fib());
}
`
	NewCompiler().Compiler(script)
}
