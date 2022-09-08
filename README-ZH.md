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
- [x] class声明。
- [x] 函数声明与调用。
- [x] 基本类型: `int/string/float/bool`
- [x] array数组类型。
- [x] 特殊类型 `nil`。
- [x] 函数类型。
- [x] 闭包：函数一等公民。
- [x] 内置函数: `len()/hash()/assertEqual()`
- [x] 标准库：`Map/LinkedList/Array`
- [x] 运算符重载
- [ ] 原生支持 `json` 支持。
- [ ] 原生 `http` 包支持。

# 例子

## Hello world
```js
println("hello world");
```

## 打印斐波那契数列

```js
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

func int() f = fib();

for (int i = 0; i < 10; i++){
    println(f());
}
```

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
```

## 数组
数组声明语法：`('[' DECIMAL_LITERAL ']')? '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'`
```js
// 声明并初始化
int[] a={1,2,3};
println(a);

// 声明一个空数组并指定大小
int[] table = [4]{};

println();
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
int[] a={1,2,3};
// len 返回数组大小
println(len(a));

// 向数组追加数据
a = append(a,4);
println(a);
// output: [1,2,3,4]

// 断言函数，不相等时会抛出运行时异常，并中断程序。
assertEqual(len(a),4);

// 返回 hashcode
int hashcode = hash(key);
```

## MapString

键值对都为 `string` 的 `HashMap`。

```js
int count =100;
MapString m1 = MapString();
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


## 联系作者


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 