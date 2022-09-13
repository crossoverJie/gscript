package gscript

import (
	"fmt"
	"testing"
)

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
func TestCompiler_Return5(t *testing.T) {
	script := `
int run(){
	for(int i=0;i<3;i++){
		for(int i=0;i<2;i++){
			//println("inner i=" +i);
			// todo crossoverJie 应该直接返回 
			//println("return finish i=" + i);
			return 100;
		}
		println("outer i=" +i);
	}
	println("不会执行");
	return 200;
}
int ret = run();
println(ret);
assertEqual(ret,100);
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Return6(t *testing.T) {
	script := `
	int run(int ret){
		int a=1;
		for (a<=3) {
			println("外层if,a=" + a);
			a++;
			return ret;
			println("return 之后不打印");
		}
		println("!!最终a="+ a);		
	}
	int ret = run(100);
	println("run 之后 ret="+ ret);
	assertEqual(ret,100);
	println("开始调第二次");
	ret = run(200);
	println("run 之后 ret="+ ret);
	assertEqual(ret,200);
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
func TestCompiler_Return4(t *testing.T) {
	script := `
int run(){
	for(1<2){
		for(1<2){
			return 100;
		}
		println(12);
		
	}
	println("不会执行");
	return 200;
}
int x=run();
println(x);
assertEqual(x,100);
int a =9999;
println(a);
assertEqual(a,9999);
`
	NewCompiler().CompilerWithoutNative(script)
}

func TestFor99(t *testing.T) {
	a := 1
	for a <= 3 {
		//if a<=2 {
		//	println("return a=",a);
		//	return;
		//}
		println("外层if,a=", a)
		a++
		return
	}
	fmt.Println("!!最终a=", a)
}

func TestFor100(t *testing.T) {

	for i := 0; i < 3; i++ {
		for i := 0; i < 2; i++ {
			println("inner i=", i)
			// todo crossoverJie 应该直接返回
			println("return finish")
			return
		}
		println("outer i=", i)
	}
}

func TestScope(t *testing.T) {
	scope()
	fmt.Println("test")
}
func scope() {
	a := 1
	for a <= 3 {
		println("外层if,a=", a)
		a++
		return
	}
	fmt.Println("!!最终a=", a)
}

func TestSlice(t *testing.T) {
	x := []int{1, 2, 3}
	if len(x) > 0 {
		x = x[:len(x)-1]
	}
	fmt.Println(x)
}

func TestTowSum(t *testing.T) {
	script := `
int twoSum(int target){
    for(1<2){
        for (3<4){
            if (5+5==10){
				println("准备return");
                return target;
            }
        }
		// 这里有注释，VisitBlockStms 特殊处理
		//println("11");
    }
    println("不会执行");
    return 100;
}
println("开始调用");
int ret = twoSum( 9);
println(ret);
assertEqual(ret,9);
`
	NewCompiler().CompilerWithoutNative(script)
}
