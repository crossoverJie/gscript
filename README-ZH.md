<div align="center">  


```
 _     _   
 ___ ___ ___ ___|_|___| |_ 
| . |_ -|  _|  _| | . |  _|
|_  |___|___|_| |_|  _|_|  
|___|             |_|   

```

📘[特性](#特性) | 🌰[例子](#例子) | 👾[REPL](#repl) | 🎉[语法](#语法) | 🎁[标准库](#标准库) | 🔧[安装](https://github.com/crossoverJie/gscript/releases) | 💡[联系作者](#联系作者) | 🇦🇺[英文文档](https://github.com/crossoverjie/gscript/blob/master/README.md)


</div><br>


# 介绍

这是一门用 Go 编写的一款**静态、强类型**的脚本语言，大部分语法参考了 `Java` 以及少量的 `Go`。

> 当前版本仅供学习与实验。

运行：

hello_world.gs:
```js
println("hello world");
```

```shell
❯ gscript hello_world.gs
hello world
```

# 特性
- [x] [class声明](#class)
- [x] [函数声明与调用](#函数)
- [x] [基本类型](#基本类型): `int/string/float/bool`
- [x] [array数组类型](#数组)
- [x] `any` [通用类型](#any)
- [x] 特殊类型 `nil`
- [x] 函数类型
- [x] [闭包：函数一等公民](#闭包)
- [x] [内置函数](#内置函数): `len()/hash()/assertEqual()/JSON()/JSONGet()`
- [x] [标准库](#标准库)
	- [x] [Map](#map)
- [x] [可变参数](#可变参数)
- [x] [运算符重载](#运算符重载)
- [x] [原生 `json` 支持](#内置函数)
- [x] [原生 `http` 包支持](#http)
- [x] 案例
	- [x] LeetCode
		- [x] [判断链表是否有环 ](https://github.com/crossoverJie/gscript/blob/main/example/linked_list_cycle141.gs)
		- [x] [两数之和](https://github.com/crossoverJie/gscript/blob/main/example/leetcode/two_sum.gs)
	- [x] [打印斐波那契数列 ](#打印斐波那契数列)
	- [x] [打印杨辉三角 ](#杨辉三角)
	- [x] [HTTP Service](https://github.com/crossoverJie/gscript/blob/main/example/http_server.gs)
-  [ ] 包管理
-  [ ] 单测命令行工具

# 例子

## Hello world
```js
println("hello world");
```

## 打印斐波那契数列

```js
func int() fib(){
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

func int() f = fib();

for (int i = 0; i < 10; i++){
    println(f());
}
```

## 网站应用

这是用 `GScript` 编写的动态网页：

[https://gscript.crossoverjie.top/index](https://gscript.crossoverjie.top/index)

源码地址：
[https://github.com/crossoverjie/gscript-homepage](https://github.com/crossoverjie/gscript-homepage)

## 杨辉三角

```js
int num(int x,int y){
	if (y==1 || y ==x) {
		return 1;
	}
    int v1 = num(x - 1, y - 1);
    int v2 = num(x - 1, y);
	int c = v1+v2;
    // int c = num(x - 1, y - 1)+num(x - 1, y);
	return c;
}

printTriangle(int row){
	for (int i = 1; i <= row; i++) {
        for (int j = 1; j <= row - i; j++) {
           print(" ");
        }
        for (int j = 1; j <= i; j++) {
            print(num(i, j) + " ");
        }
        println("");
    }
}

printTriangle(7);

// output:
      1 
     1 1 
    1 2 1 
   1 3 3 1 
  1 4 6 4 1 
 1 5 10 10 5 1 
1 6 15 20 15 6 1
```

---

更多例子：[https://github.com/crossoverJie/gscript/tree/main/example](https://github.com/crossoverJie/gscript/tree/main/example)

# REPL
```shell
> ./gscript 
```

![](doc/repl.gif)

# 语法

## 基本类型

当前版本支持 `int/string/float/bool` 四种基本类型以及 `nil` 特殊类型。

变量声明语法：`type identifier (= expr)?`。

```js
int a=10;
string b,c;
float e = 10.1;
bool f = false;
string x = ^
{
    "name": "bob",
    "age": 20,
    "skill": {
        "lang": [
            {
                "go": {
                    "feature": [
                        "goroutine",
                        true
                    ]
                }
            }
        ]
    }
}^;
```

## 数组
数组声明语法：`('[' DECIMAL_LITERAL ']')? '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'`
```js
// 声明并初始化
int[] a={1,2,3};
println(a);

// 声明一个空数组并指定大小
int[] table = [4]{};

println("");
// 向数组 append 数据
a = append(a,4);
println(a);
for(int i=0;i<len(a);i++){
	println(a[i]);
}

// 通过下标获取数组数据
int b=a[2];
println(b);
```

## any

`any` 是通用类型，类似于 Java 中的 Object 和 Go 中的 `interface{}`。

```js
any a =10;
println(a);

int fun1(any a,int b){
	return a+b;
}
int v =fun1(1,2);
println(v);
assertEqual(v,3);

any v2 = fun1(1,2);
println(v2);
assertEqual(v2,3);

int fun2(int a, any b){
	return a+b;
}
int v3 =fun2(1,2);
println(v3);
assertEqual(v3,3);

int fun3(any a, any b){
	return a+b;
}
int v4 =fun3(1,2);
println(v4);
assertEqual(v4,3);

int fun4(any a, any b){
	return a+b;
}
string v5 =fun4("10", "20");
println(v5);
assertEqual(v5,"1020");
```

[标准库](https://github.com/crossoverJie/gscript/blob/main/internal/internal.gs#L25)中有相关应用。
## Class

自定义 Class 与 Java 类似：

```js
class ListNode{
    int value;
    ListNode next;
    ListNode(int v, ListNode n){
        value =v;
        next = n;
    }
}

// 调用构造函数时不需要使用 new 关键字。
ListNode l1 = ListNode(1, nil);

// 使用 . 调用对象属性或函数。
println(l1.value);
```

缺省情况下 `class` 具有无参构造函数：

```js
class Person{
	int age=10;
	string name="abc";
	int getAge(){
		return 100+age;
	}
}

// 无参构造函数
Person xx= Person();
println(xx.age);
assertEqual(xx.age, 10);
println(xx.getAge());
assertEqual(xx.getAge(), 110);
```


## 函数

```js
// 判断链表是否有环
bool hasCycle(ListNode head){
    if (head == nil){
        return false;
    }
    if (head.next == nil){
        return false;
    }

    ListNode fast = head.next;
    ListNode slow = head;
    bool ret = false;
    for (fast.next != nil){
        if (fast.next == nil){
            return false;
        }
        if (fast.next.next == nil){
            return false;
        }
        if (slow.next == nil){
            return false;
        }
        if (fast == slow){
            ret = true;
            return true;
        }

        fast = fast.next.next;
        slow = slow.next;
    }
    return ret;
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
assertEqual(b3, true);
```

函数声明语法：`typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')*`

```js
add(int a){}
```

> 当函数没有返回值时，可以声明为 void 或直接忽略返回类型。


## 闭包

在 `GScript` 中，函数作为一等公民可以作为变量传递，同时也能实现闭包。

函数类型语法：`func typeTypeOrVoid '(' typeList? ')'`

```js
// 外部变量，全局共享。
int varExternal =10;
func int(int) f1(){
	// 闭包变量对每个闭包单独可见
	int varInner = 20;
	int innerFun(int a){
		println(a);
		int c=100;
		varExternal++;
		varInner++;
		return varInner;
	}
	return innerFun;
}

// f2 作为一个函数类型，接收的是一个返回值和参数都是 int 的函数。
func int(int) f2 = f1();
for(int i=0;i<2;i++){
	println("varInner=" + f2(i) + ", varExternal=" + varExternal);
}
println("=======");
func int(int) f3 = f1();
for(int i=0;i<2;i++){
	println("varInner=" + f3(i) + ", varExternal=" + varExternal);
}
```

最终输出如下：

```shell
0
varInner=21, varExternal=11
1
varInner=22, varExternal=12
=======
0
varInner=21, varExternal=13
1
varInner=22, varExternal=14

```

## 可变参数

`GScript` 支持函数定义可变参数，语法如下：

```js
int add(string s, int ...num){
	println(s);
	int sum = 0;
	for(int i=0;i<len(num);i++){
		int v = num[i];
		sum = sum+v;
	}
	return sum;
}
int x = add("abc", 1,2,3,4);
println(x);
assertEqual(x, 10);
```


## 运算符重载
`GScript` 支持以下运算符重载：
- `+-*/`
- `== != < <= > >=`

> 重载函数名称必须是 **operator**，名称后跟上运算符即可重载。

```js
class Person{
	int age;
	Person(int a){
		age = a;
	}
}
Person operator + (Person p1, Person p2){
	Person pp = Person(p1.age+p2.age);
	return pp;
}
Person operator - (Person p1, Person p2){
	Person pp = Person(p1.age-p2.age);
	return pp;
}
Person operator * (Person p1, Person p2){
	Person pp = Person(p1.age * p2.age);
	return pp;
}
Person operator / (Person p1, Person p2){
	Person pp = Person(p1.age / p2.age);
	return pp;
}
bool operator == (Person p1, Person p2){
	return p1.age==p2.age;
}
bool operator != (Person p1, Person p2){
	return p1.age!=p2.age;
}
bool operator > (Person p1, Person p2){
	return p1.age>p2.age;
}
bool operator >= (Person p1, Person p2){
	return p1.age>=p2.age;
}
bool operator < (Person p1, Person p2){
	return p1.age<p2.age;
}
bool operator <= (Person p1, Person p2){
	return p1.age<=p2.age;
}
Person p1 = Person(10);
Person p2 = Person(20);
//Person p3 =  operator(p1,p2);
Person p3 = p1+p2;
println("p3.age="+p3.age);
assertEqual(p3.age, 30);

Person p4 = p1-p2;
println("p4.age="+p4.age);
println(100-10);

Person p5 = p1*p2;
println("p5.age="+p5.age);
assertEqual(p5.age, 200);

Person p6 = p2/p1;
println("p6.age="+p6.age);
assertEqual(p6.age, 2);

bool b1 = p1 == p2;
println("b1=="+b1);
assertEqual(b1,false);

bool b2 = p1 != p2;
println("b2=="+b2);
assertEqual(b2,true);

bool b3 = p1 > p2;
println("b3=="+b3);
assertEqual(b3,false);

bool b4 = p1 >= p2;
println("b4=="+b4);
assertEqual(b4,false);

bool b5 = p1 < p2;
println("b5=="+b5);
assertEqual(b5,true);

bool b6 = p1 <= p2;
println("b6=="+b6);
assertEqual(b6,true);
```


更多样例请参考：[https://github.com/crossoverJie/gscript/tree/main/example](https://github.com/crossoverJie/gscript/tree/main/example)

# 标准库

标准库源码：[https://github.com/crossoverJie/gscript/tree/main/internal](https://github.com/crossoverJie/gscript/tree/main/internal)

## 内置函数

```js

printf("hello %s ","123");
printf("hello-%s-%s ","123","abc");
printf("hello-%s-%d ","123",123);

string s = sprintf("nice to meet %s", "you");
println(s);
assertEqual(s,"nice to meet you");

int[] a={1,2,3};
// len 返回数组大小
println(len(a));

// 向数组追加数据
append(a,4);
println(a);
// output: [1,2,3,4]

// 断言函数，不相等时会抛出运行时异常，并中断程序。
assertEqual(len(a),4);

// 返回 hashcode
int hashcode = hash(key);

// 序列化 JSON 字符串
class P{
	string name;
	P(string n){
		name = n;
	}
}
class Object{
	P p;
	int x;
	Object(P pp, int xx){
		p = pp;
		x = xx;
	}
}
P p1 = P("abc");
Object o1 = Object(p1, 100);
string json = JSON(o1);
println(json); //{"p":{"name":"abc"},"x":100}

// 查询 JSON
int x = JSONGet(json,"x");
println(x);
assertEqual(x,100);

string name = JSONGet(json,"p.name");
println(name);
assertEqual(name,"abc");

// 获取启动参数
System s = System();
string[] args = s.getOSArgs();
```



> 更多 JSON 查询语法请参考：[xjson](https://github.com/crossoverJie/xjson#arithmetic-syntax)

## Map
函数定义：
```js
class Map{
	put(any key, any value){}
	any get(any key){}
}
```

```js
int count =100;
Map m1 = Map();
for (int i=0;i<count;i++){
	string key = i+"";
	string value = key;
	m1.put(key,value);
}
println(m1.getSize());
assertEqual(m1.getSize(),count);

for (int i=0;i<count;i++){
	string key = i+"";
	string value = m1.get(key);
	println("key="+key+ ":"+ value);
	assertEqual(key,value);
}
```

# http
标准库定义：

```js
// http lib
// Response json
FprintfJSON(int code, string path, string json){}
// Resonse html
FprintfHTML(int code, string path, string html){}

// path (relative paths may omit leading slash)
string QueryPath(string path){}

string FormValue(string path, string key){}
class HttpContext{
    string path;
    JSON(int code, any v){
        string json = JSON(v);
        FprintfJSON(code, path, json);
    }
    HTML(int code, any v) {
        string html = v;
        FprintfHTML(code, path, html);
    }
    string queryPath() {
        string p = QueryPath(path);
        return p;
    }

    string formValue(string key){
        string v = FormValue(path, key);
        return v;
    }
}
// Bind route
httpHandle(string method, string path, func (HttpContext) handle){
    // println("path="+path);
    HttpContext ctx = HttpContext();
    handle(ctx);
}
// Run http server.
httpRun(string addr){}
```

启动一个 `HTTP` 服务：

```js
class Person{
    string name;
}
func (HttpContext) handle (HttpContext ctx){
    Person p = Person();
    p.name = "abc";
    println("p.name=" + p.name);
    println("ctx=" + ctx);
    ctx.JSON(200, p);
}

func (HttpContext) handle1 (HttpContext ctx){
    Person p = Person();
    p.name = "def";
    println("p.name=" + p.name);
    println("ctx=" + ctx);
    ctx.JSON(200, p);
}
func (HttpContext) handle2 (HttpContext ctx){
    DateTime d = DateTime();
    string local = d.getCurrentTime("Asia/Shanghai","2006-01-02 15:04:05");
    println(local);
    string html =^
    <html>
    <title>hello</title>
    <h1>current ^+ local +^</h1>
    <p>hahaha</p>
    </html>
    ^;
    string queryPath = ctx.queryPath();
    println("queryPath = " + queryPath);

    // http://127.0.0.1:8000/p/2?id=100
    string id = ctx.formValue("id");
    println("id="+id);
    ctx.HTML(200, html);
}
httpHandle("get", "/p", handle);
httpHandle("get", "/p/1", handle1);
httpHandle("get", "/p/2", handle2);
httpRun(":8000");

```



## 联系作者


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 