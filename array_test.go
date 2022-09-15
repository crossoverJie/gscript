package gscript

import (
	"os"
	"testing"
)

func Test_array(t *testing.T) {
	script := `
int[] a={1,2,3};
// [1 2 3]
println(a);
a[0]=2;
// [2 2 3]
println(a);
a[1]=3;
// [2 3 3]
println(a);
a[2]=4;
// [2 3 4]
println(a);
`
	NewCompiler().Compiler(script)
}
func Test_array2(t *testing.T) {
	script := `
int[] a={1,2,3};
for(int i=0;i<3;i++){
	println(a[i]);
}
int b=a[2];
println(b);
assertEqual(b,3);
`
	NewCompiler().Compiler(script)
}
func Test_array4(t *testing.T) {
	script := `
int[] a={1,2,3};
for(int i=0;i<3;i++){
	println(a[i]);
	int temp = a[i];
	a[i]=temp+1;
}
println();
for(int i=0;i<3;i++){
	println(a[i]);
	int ret = a[i];
	assertEqual(ret, i+2);
}
`
	NewCompiler().Compiler(script)
}
func Test_array5(t *testing.T) {
	script := `
int[] a={1,2,3};
println(a);
println();
a = append(a,4);
println(a);
for(int i=0;i<len(a);i++){
	println(a[i]);
	int temp = a[i];
	assertEqual(temp, i+1);
}
`
	NewCompiler().Compiler(script)
}
func Test_array6(t *testing.T) {
	script := `
int[] a={1,2,3};
println(a);
assertEqual(len(a),3);
println();
a = append(a,4);
println(a);
for(int i=0;i<len(a);i++){
	println(a[i]);
	int temp = a[i];
	assertEqual(temp, i+1);
}
assertEqual(len(a),4);
`
	NewCompiler().Compiler(script)
}
func Test_array7(t *testing.T) {
	script := `
string[] a={"1","2","3"};
int[] b={1,2,3};
if (3==len(a)){
	println("3==len(a)");
}
for(int i=0;i<len(a);i++){
	println(a[i]);
}
println("======");
for(int i=0;i<len(b);i++){
	println(b[i]);
}
println("======");
bool[] c ={true,true,false};
for(int i=0;i<len(c);i++){
	println(c[i]);
}
println("======");
float[] d ={10.1,1.1,2.1};
for(int i=0;i<len(d);i++){
	println(d[i]);
}
`
	NewCompiler().Compiler(script)
}
func Test_array8(t *testing.T) {
	script := `
int[] a=[10]{};
println(len(a));
assertEqual(len(a), 10);
a = append(a,1);
println(len(a));
assertEqual(len(a), 11);
`
	NewCompiler().Compiler(script)
}
func Test_array9(t *testing.T) {
	script := `
int[] a=[2]{};
println("数组大小:"+len(a));
assertEqual(len(a), 2);
a = append(a,1);
println("数组大小:"+len(a));
assertEqual(len(a), 3);
println(a);
a[0]=100;
println(a);
int temp = a[0];
assertEqual(temp, 100);
`
	NewCompiler().Compiler(script)
}
func Test_array10(t *testing.T) {
	script := `
class Person{
	int age;
	string name;
	Person(int a, string n){
		age = a;
		name =n;
	}
}
Person p1 = Person(1,"a");
println(p1);
Person p2 = Person(1,"b");
Person[] list = {p1,p2};
//[{a 100} {a 10}]
// [[1 a] [1 b]]
println(list);
//println(list[0]);
Person temp = list[0];
println(temp.name);
assertEqual(temp.name,"a");
assertEqual(temp.age,1);
`
	NewCompiler().Compiler(script)
}
func Test_array11(t *testing.T) {
	script := `
	string[] args = {"1","2","3"};
    string port = args[2];
    println(":" +args[2]);
`
	NewCompiler().Compiler(script)
}
func Test_array12(t *testing.T) {
	script := `
	int a=10;
	int[] x =[2]{};
	x[0]=a;
	println(x);
`
	NewCompiler().Compiler(script)
}
func Test_array13(t *testing.T) {
	script := `
int[] n ={1,2,3};
int l = n[0];
int r = n[1];
if(l+r == 3){
	println(3);
}
assertEqual(l+r == 3, true);

if( n[0] + n[1] == 3){
	println(3);
}
assertEqual(n[0] + n[1] == 3, true);

int x = n[0] + n[1];
println(x);
assertEqual(x,3);

`
	NewCompiler().Compiler(script)
}
func Test_OSArgs14(t *testing.T) {
	Args = os.Args
	script := `
string[] args = getOSArgs();
println(len(args));
println(args);
`
	NewCompiler().Compiler(script)
}

func Test_array14(t *testing.T) {
	script := `
int[] a={1,2,3};
println(a);
println();
append(a,4);
int b = a[3];
assertEqual(4, b);
println(a);
for(int i=0;i<len(a);i++){
	println(a[i]);
	int temp = a[i];
	assertEqual(temp, i+1);
}
`
	NewCompiler().Compiler(script)
}
