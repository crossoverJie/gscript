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
	NewCompiler().Compiler(script)
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
			println("return finish")
			return
		}
		println("outer i=", i)
	}
}
func TestFor101(t *testing.T) {
	script := `
	for (int i = 0; i < 3; i++) {
		for (int i=0; i < 2; i++) {
			println("inner i="+i);
			println("return finish");
			return;
		}
		println("outer i="+i);
	}
`
	NewCompiler().Compiler(script)
}

func TestFor98(t *testing.T) {
	script := `
int f1(int x){
	f1(1);
}
`
	NewCompiler().CompilerWithoutNative(script)
}

func TestFor102(t *testing.T) {
	script := `
	printf(string format, any ...a){}
	println(any a){}
	assertEqual(any a1, any a2){}
	int r(int x, int ret){
		if(x==0){
			return ret;
		}else{
			ret = ret+1;
		}
		int b = r(x-1, ret);
		printf("b=%d ", b);
		return b;
	}
	
	int i = r(10, 0);
	println(i);
	assertEqual(i,10);
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestFor103(t *testing.T) {
	script := `
	printf(string format, any ...a){}
	println(any a){}
	string print(any ...a){}
	//assertEqual(any a1, any a2){}
	int r(int x, int ret){
		if(x==0){
			return ret;
		}else{
			ret = ret+1;
		}
		int b = r(x-1, ret);
		printf("b=%d ", b);
		return b;
	}

	int num(int x,int y){
		if (y==1 || y ==x) {
			return 1;
		}
		int c = num(x - 1, y - 1) + num(x - 1, y);
		printf("c=%d ", c);
		return c;
	}

	fun(int row){
		for (int i = 1; i <= row; i++) {
			for (int j = 1; j <= i; j++) {
            	int i = r(10, 0);
				//println(i);
				//print(num(i, j) + " ");
        	}
		}
	}
	
	fun(10);
	//assertEqual(i,10);
`
	NewCompiler().CompilerWithoutNative(script)
}

func TestRe(t *testing.T) {
	i := r(10, 0)
	fmt.Println(i)
}

func r(x, ret int) int {

	if x == 0 {
		return ret
	} else {
		ret = ret + 1
	}
	i := r(x-1, ret)
	fmt.Println("i=" + fmt.Sprint(i))
	return i
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
	print(5)
	//fmt.Print("123" + " ")
	//fmt.Print("456" + " ")
}
func num(x, y int) int {
	if y == 1 || y == x {
		return 1
	}
	c := num(x-1, y-1) + num(x-1, y)
	return c
}
func print(row int) {
	for i := 0; i < row; i++ {
		for j := i; j <= row-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			v := num(i, j)
			fmt.Print(fmt.Sprint(v) + " ")
		}
		fmt.Println("")
	}
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
	NewCompiler().Compiler(script)
}
