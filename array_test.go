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
