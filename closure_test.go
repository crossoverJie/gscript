package gscript

import "testing"

func TestClosure(t *testing.T) {
	script := `class ListNode{
    int value;
    ListNode next;
    ListNode(int v, ListNode n){
        value =v;
        next = n;
    }
}

// 两个对象比较需要实现运算符重载
bool operator == (ListNode p1, ListNode p2){
    return p1.value == p2.value;
}

bool hasCycle(ListNode head){
    if (head == nil){
        return false;
    }
    if (head.next == nil){
        return false;
    }

    ListNode fast = head.next;
    ListNode slow = head;
    // bool ret = false;
    for (fast.next != nil){

        if (fast == slow){
            return true;
        }

        if (fast.next == nil){
            return false;
        }
        if (fast.next.next == nil){
            return false;
        }
        if (slow.next == nil){
            return false;
        }

        fast = fast.next.next;
        slow = slow.next;
    }
    return false;
}

ListNode l1 = ListNode(1, nil);
bool b1 =hasCycle(l1);
println(b1);
assertEqual(b1, false);

ListNode l4 = ListNode(4, nil);
ListNode l3 = ListNode(3, l4);
ListNode l2 = ListNode(2, l3);
bool b2 = hasCycle(l2);
println(b2);
assertEqual(b2, false);

l4.next = l2;
bool b3 = hasCycle(l2);
println(b3);
assertEqual(b3, true);`
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
//assertEqual(f2(), 20);
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
func int() fun(){
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
func TestClosure8(t *testing.T) {
	script := `
// 函数类型的变量
func (int, int) handle1 (int a, int b){
	println("a="+a + " b="+b);
}
//handle1(10,20);
//println("====");

handleFunc(string path, func (int, int) handle){
	println("path="+path);
	handle(100,200);
}
func (int, int) x = handle1;
handleFunc("/abc", handle1);
`
	NewCompiler().Compiler(script)
}
