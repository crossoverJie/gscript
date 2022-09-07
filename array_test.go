package gscript

import "testing"

func Test_array(t *testing.T) {
	script := `
int[] a={1,2,3};
println(a);
a[0]=2;
println(a);
a[1]=3;
println(a);
a[2]=4;
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
	int temp = a[i]
	a[i]=temp+1;
}
println();
for(int i=0;i<3;i++){
	println(a[i]);
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
for(int i=0;i<4;i++){
	println(a[i]);
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
a = append(a,1);
println(len(a));
`
	NewCompiler().Compiler(script)
}
func Test_array9(t *testing.T) {
	script := `
int[] a=[2]{};
println("数组大小:"+len(a));
a = append(a,1);
println("数组大小:"+len(a));
println(a);
a[0]=100;
println(a);
`
	NewCompiler().Compiler(script)
}
