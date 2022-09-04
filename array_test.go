package gscript

import "testing"

func Test_array(t *testing.T) {
	script := `
int[] a={1,2,3};
print(a);
a[0]=2;
print(a);
a[1]=3;
print(a);
a[2]=4;
print(a);
`
	NewCompiler().Compiler(script)
}
func Test_array2(t *testing.T) {
	script := `
int[] a={1,2,3};
for(int i=0;i<3;i++){
	print(a[i]);
}
int b=a[2];
print(b);
assertEqual(b,3);
`
	NewCompiler().Compiler(script)
}
func Test_array4(t *testing.T) {
	script := `
int[] a={1,2,3};
for(int i=0;i<3;i++){
	print(a[i]);
	int temp = a[i]
	a[i]=temp+1;
}
print();
for(int i=0;i<3;i++){
	print(a[i]);
}
`
	NewCompiler().Compiler(script)
}
func Test_array5(t *testing.T) {
	script := `
int[] a={1,2,3};
print(a);
print();
a = append(a,4);
print(a);
for(int i=0;i<4;i++){
	print(a[i]);
}
`
	NewCompiler().Compiler(script)
}
func Test_array6(t *testing.T) {
	script := `
int[] a={1,2,3};
print(a);
assertEqual(len(a),3);
print();
a = append(a,4);
print(a);
for(int i=0;i<len(a);i++){
	print(a[i]);
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
	print("3==len(a)");
}
for(int i=0;i<len(a);i++){
	print(a[i]);
}
print("======");
for(int i=0;i<len(b);i++){
	print(b[i]);
}
print("======");
bool[] c ={true,true,false};
for(int i=0;i<len(c);i++){
	print(c[i]);
}
print("======");
float[] d ={10.1,1.1,2.1};
for(int i=0;i<len(d);i++){
	print(d[i]);
}
`
	NewCompiler().Compiler(script)
}
func Test_array8(t *testing.T) {
	script := `
int[] a=[10]{};
print(len(a));
a = append(a,1);
print(len(a));
`
	NewCompiler().Compiler(script)
}
