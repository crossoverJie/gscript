package gscript

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompiler_Compiler(t *testing.T) {
	script := `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) ;
} else {
	return 20 ;
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Compiler2(t *testing.T) {
	script := `
int a=10;
a++;
return a;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
	assert.Equal(t, object.GetReturnObject(), 11)
}
func TestCompiler_CompilerFor(t *testing.T) {
	script := `
int age = 0 ;
for(int i = 0;i<100;i++) {
	age++;
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
	assert.Equal(t, object.GetReturnObject(), 100)
}
func TestCompiler_CompilerFor3(t *testing.T) {
	script := `
int age = 0 ;
int sum=1;
for(int i = 0;i<100;i++) {
	sum=sum+1;
	for(int i = 0;i<100;i++) {
		age++;
	}
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
	assert.Equal(t, object.GetReturnObject(), 100*100)
}
func TestCompiler_CompilerFor2(t *testing.T) {
	script := `
int age = 1 ;
for(int i = 0;i<100;i++) {
	age=age+1;
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
	assert.Equal(t, object.GetReturnObject(), 101)
}
func TestCompiler_CompilerIf(t *testing.T) {
	script := `
int age=10;
if (age>10){
	age++;
}else{
	age--;
}
return age;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
	assert.Equal(t, object.GetReturnObject(), 9)
}
func TestCompiler_CompilerIf2(t *testing.T) {
	script := `
int age=10;
if (age>10){
	age++;
}else{
	age--;
}
return abc;
`
	compiler := NewCompiler().Compiler(script)
	object := compiler.(*ReturnObject)
	fmt.Println(object.GetReturnObject())
}

func TestCompiler_FunctionCall(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a) {
	return a+b+3;
}
myfunc(2);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 15)
}
func TestCompiler_FunctionCall2(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a,int d) {
	return a+b+3+d;
}
myfunc(2,10);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 25)
}
func TestCompiler_FunctionCall3(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a,int b) {
	return a+b+3;
}
myfunc(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 25)
}
func TestCompiler_FunctionCall4(t *testing.T) {
	script := `
int b= 10;
int foo(){
	return b;
}
int myfunc(int a,int b) {
	int e = foo();
	return a+b+3+e;
}
myfunc(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 35)
}
func TestCompiler_FunctionCall5(t *testing.T) {
	script := `
int b= 10;
int foo(int age){
	for(int i=0;i<10;i++){
		age++;
	}
	return b+age;
}
int add(int a,int b) {
	int e = foo(10);
	e = e+10;
	return a+b+3+e;
}
add(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 65)
}
func TestCompiler_FunctionCall7(t *testing.T) {
	script := `
int a=10;
int b=30;
int fun(int f){
	return f+100+b;
}
fun(a);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 140)
}
func TestCompiler_FunctionCall8(t *testing.T) {
	script := `
int a=10;
a;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	//assertEqual.Equal(t, compiler, 55)
}

//func TestCompiler_FunctionCall6(t *testing.T) {
//	script := `
//int b= 10;
//int,int foo(int age){
//	return b+age,100;
//}
//int myfunc(int a,int b) {
//	int e,f = foo(10);
//	e = e+10;
//	return a+b+3+e;
//}
//myfunc(2,20);
//`
//	compiler := NewCompiler().Compiler(script)
//	fmt.Println(compiler)
//	assertEqual.Equal(t, compiler, 55)
//}

func TestCompiler_Class(t *testing.T) {
	script := `
class Person{
	int age=10;
	string name="abc";
	int getAge(){
		return 100+age;
	}
	int sub(){
		return age-10;
	}
	int div(){
		return age/10;
	}
	int mul(){
		return age*10;
	}
	int mod(){
		return age%10;
	}
	string getName(){
		return age+name;
	}
}
Person xx= Person();
println(xx.age);
assertEqual(xx.age, 10);
println(xx.getAge());
int r1 = xx.age;
int r2 = xx.getAge();
assertEqual(r1,10);
assertEqual(r2,110);
xx.age=200;
println(xx.age);
assertEqual(xx.age,200);
println(xx.getName());
assertEqual(xx.getName(),"200abc");
assertEqual(xx.sub(),190);
assertEqual(xx.div(),20);
assertEqual(xx.mul(),2000);
assertEqual(xx.mod(),0);

if(10>2){
	println("10>2");
}
if(2<xx.age){
	println("2<xx.age");
}
if(2<=2){
	println("2<=2");
}
if(2>=2){
	println("2>=2");
}
if(200==xx.age){
	println("200==xx.age");
}
if(xx.getName()=="200abc"){
	println("xx.getName()==200abc");
}
if(xx.getName()!="200abc1"){
	println("xx.getName()!=200abc1");
}

`
	NewCompiler().Compiler(script)
}

func TestCompiler_Class2(t *testing.T) {
	script := `
class Person{
	int age=0;
	Person(int a){
		age = a;
	}
}
Person xx= Person(10);
println(xx.age);
assertEqual(10,xx.age);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Class3(t *testing.T) {
	script := `
class Person{
	int age=0;
	int number=0;
	Person(int a,int n){
		age = a;
		number = n;
	}
}
Person xx= Person(10,20);
println(xx.age);
println(xx.number);
assertEqual(10,xx.age);
assertEqual(20,xx.number);
int age=0;
assertEqual(age+10,10);
assertEqual(10+1.1, 11.1);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_For(t *testing.T) {
	script := `
int a=0;
for(a<=10){
	println(a);
	a++;
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_For2(t *testing.T) {
	script := `
int a=0;
for(a<=10){
	println(a);
	a++;
	if(a==5){
		break;
	}
}
assertEqual(a,5);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_For0(t *testing.T) {
	script := `
int a=0;
for(a<3){
	a++;
	if(a==2){
		continue;
	}
	println(a);
}
assertEqual(a,3);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_For3(t *testing.T) {
	script := `
for(int i=0;i<5;i++){
	println(i);
	if (i==3){
		break;
	}
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_For4(t *testing.T) {
	script := `
for(int i=0;i<5;i++){
	if (i==3){
		continue;
	}
	println(i);
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Return(t *testing.T) {
	script := `
for(int i=0;i<5;i++){
	if (i==3){
		return;
	}
	println(i);
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Return2(t *testing.T) {
	script := `
int a=0;
for(a<=10){
	println(a);
	a++;
	if(a==5){
		return;
	}
}
assertEqual(a,5);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Return3(t *testing.T) {
	script := `
int a=0;
bool x = false;
for(a<=10){
	a++;
	if(a==5){
		x =true;
		return;
	}
	println("after return");
}
if(x){
	println(x);
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Nil(t *testing.T) {
	script := `
class Test{
}
t(Test t){
	println("t");
}
Test x = Test();
t(nil);

t2(int x){
	println("t2");
}
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestCompiler_Condition(t *testing.T) {
	script := `
bool a = true;
bool b = true;
if (a&&b){
	println("a&&b");
}
b = false;
if (!(a&&b)){
	println("!a&&b");
}
int c =100;
int d =100;
if(c==100 && d==100){
	println("a==c");
}

if(a||b){
	println("a||b");
}
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestCompiler_Main(t *testing.T) {
	script := `
main();
class Test{
	p(){
		println("123");
	}
}
main(){
	MapString m1 = MapString();
	int count =3;
	for (int i=0;i<count;i++){
		string key = i+"";
		string value = key;
		m1.put(key,value);
	}
	for (int i=0;i<count;i++){
		string key = i+"";
		string value = m1.get(key);
		println("key="+key+ ":"+ value);
		assertEqual(key,value);
	}
	Test t = Test();
	t.p();
}
println("999")
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Class_EQ(t *testing.T) {
	script := `
class Test{
	int value;
	Test(int v){
		value =v;
	}
}
Test t1 = Test(1);
Test t2 = t1;
println(t1==t2);
`
	NewCompiler().Compiler(script)
}

func TestNativeReturn(t *testing.T) {
	s := testReturn()
	println(s)
}

func testReturn() string {
	a := 0
	for a <= 10 {
		a++
		if a == 5 {
			return "1"
			fmt.Println("return")
		}
		fmt.Println("after return")
	}
	return "0"
}
