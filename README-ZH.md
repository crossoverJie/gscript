
```
 _     _   
 ___ ___ ___ ___|_|___| |_ 
| . |_ -|  _|  _| | . |  _|
|_  |___|___|_| |_|  _|_|  
|___|             |_|   

```

<div align="center">  

📘[特性](#特性) 🌰[例子](#例子) 🎉[语法](#语法) 🔧[安装](https://github.com/crossoverJie/gscript/releases)💡[联系作者](#联系作者)| 🇦🇺[英文文档](https://github.com/crossoverjie/gscript/blob/master/README.md)


</div><br>


# 介绍

这是一门用 Go 编写的一款**静态、强类型**的脚本语言，大部分语法参考了 Java 以及少量的 Go。

> 当前版本仅供学习与实验。


# 特性
- [x] class声明。
- [x] 函数声明与调用。
- [x] 基本类型: `int/string/float/bool`
- [x] 特殊类型 `nil`
- [x] 函数类型。
- [x] 闭包：函数一等公民。
- [x] 内置函数: `len()/hash()/assertEqual()`
- [x] 标准库：`Map/LinkedList/Array`
- [ ] 原生支持 `json` 支持。
- [ ] 原生 `http` 包支持。

# 例子

## Hello world
```js
println("hello world");
```

## 打印斐波那契数列

```js
void fib(){
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

# 语法

## 基本类型

当前版本支持 `int/string/float/bool` 四种基本类型以及 `nil` 特殊类型。

变量声明语法：`type identifier (= expr)?`。

```js
int a=10;
string b,c;
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

更多样例请参考：[https://github.com/crossoverJie/gscript/tree/main/example](https://github.com/crossoverJie/gscript/tree/main/example)

## 联系作者


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 