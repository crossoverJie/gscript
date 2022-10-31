package gscript

import "testing"

func TestReturn(t *testing.T) {
	script := `
class T{}
T getNumber(){
	T t = T();
	return t;
}
`
	NewCompiler().Compiler(script)
}
func TestReturn2(t *testing.T) {
	script := `
class T{}
int getNumber(){
	T t = T();
	return t;
}
`
	NewCompiler().Compiler(script)
}
func TestReturn3(t *testing.T) {
	script := `
int getNumber(){
  return true;  // 返回值应该是一个整型而不是布尔
}
println(getNumber());
`
	NewCompiler().Compiler(script)
}
func TestReturn4(t *testing.T) {
	script := `
bool getNumber(){
  return true;  
}
println(getNumber());
`
	NewCompiler().Compiler(script)
}

func TestReturnTypeFail1(t *testing.T) {
	script := `
class Test{
	int value=0;
	Test(int v){
		value=v;
	}

	int map(func string(int) f){
		return f(value);
	}
}

string add(int v){
	return "abc"; 
}

Test t =Test(100);

func string(int) a=add;

println(t.map(a));
assertEqual(t.map(a),101);


`
	NewCompiler().Compiler(script)
}
