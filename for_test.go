package gscript

import "testing"

func TestFor(t *testing.T) {
	script := `
int[] nums = {1,2,3};
for (int j = len(nums) -1 ;j > 0 ;j--){
	int r = nums[j];
	println(r);
	//println(j);
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
	println("after return=" + a);
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
	println("after return="+a);
}
if(x){
	println(x);
}
`
	NewCompiler().Compiler(script)
}
