


<div align="center">  

```
 _     _   
 ___ ___ ___ ___|_|___| |_ 
| . |_ -|  _|  _| | . |  _|
|_  |___|___|_| |_|  _|_|  
|___|             |_|   

```

📘[Features](#features) | 🌰[Demo](#demo) | 👾[REPL](#repl) | 🎉[Syntax](#syntax) | 🎁[Standard library](#standard-library) | 🔧[Install](https://github.com/crossoverJie/gscript/releases) | 💡[Contact Author](#contact-author) | 🇨🇳[中文文档](https://github.com/crossoverJie/gscript/blob/main/README-ZH.md)



</div><br>

# Introduction

This is a **statically and strongly** typed language written in Go, the syntax of Java and Go is referenced.

> The current version is for study and experimentation only.

hello_world.gs:

```js
println("hello world");
```

```shell
❯ gscript hello_world.gs
hello world
```

# Features

- [x] Class declaration.
- [x] Function declaration and call.
- [x] Primitive type: `int/string/float/bool`.
- [x] Array type.
- [x] `nil` type.
- [x] `any` type.
- [x] Function type.
- [x] Closure：Functions as First-Class Objects.
- [x] Native function: `len()/hash()/assertEqual()`.
- [x] Standard library：`Map/LinkedList/Array`.
- [x] Operator overloading.
- [ ] Native support `json`.
- [ ] Native support `http`.



# Demo

## Hello world
```js
println("hello world");
```

## Print fibonacci

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

# Syntax

## Primitive

The current version supports four primitive type: `int/string/float/bool` and `nil` type.

Variable declaration syntax: `type identifier (= expr)?`.

```js
int a=10;
string b,c;
float e = 10.1;
bool f = false;
```

## Array
Array declaration syntax: `('[' DECIMAL_LITERAL ']')? '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'`

```js
// Declare and initialize
int[] a={1,2,3};
println(a);

// Declare an empty array and specify the length
int[] table = [4]{};

println();
// Append data to array.
a = append(a,4);
println(a);
for(int i=0;i<len(a);i++){
	println(a[i]);
}

// Access to data by index.
int b=a[2];
println(b);
```

## any type
An `any` type may hold values of all type.
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

Example: [Standard library](https://github.com/crossoverJie/gscript/blob/main/internal/internal.gs#L25)

## Class
```js
class ListNode{
    int value;
    ListNode next;
    ListNode(int v, ListNode n){
        value =v;
        next = n;
    }
}

// The new keyword is not required to call the constructor.
ListNode l1 = ListNode(1, nil);

// Using . to access object property or method.
println(l1.value);
```

The default comes with a parameterless constructor

```js
class Person{
	int age=10;
	string name="abc";
	int getAge(){
		return 100+age;
	}
}

// parameterless constructor
Person xx= Person();
println(xx.age);
assertEqual(xx.age, 10);
println(xx.getAge());
assertEqual(xx.getAge(), 110);
```

## function

```js
// cycle linked list
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

Function declaration syntax: `typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')*`

```js
add(int a){}
```

> When there is no return value, the return type can also be ignored.


## Closure

Function type syntax: `func typeTypeOrVoid '(' typeList? ')'`

```js
// External variable, global shared.
int varExternal =10;
func int(int) f1(){
	// Closure variable.
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

// f2 is a function type, the return type and parameter are both int.
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

Output:

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

## Operator overloading

Operator that `Gscript` support:

- `+-*/`
- `== != < <= > >=`

> Overloading function must be **operator**, and append operator.

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


More examples: [https://github.com/crossoverJie/gscript/tree/main/example](https://github.com/crossoverJie/gscript/tree/main/example)

# Standard library

Standard library source code: [https://github.com/crossoverJie/gscript/tree/main/internal](https://github.com/crossoverJie/gscript/tree/main/internal)


## Native function
```js
int[] a={1,2,3};
// len return array length.
println(len(a));

// Append data to array.
a = append(a,4);
println(a);
// output: [1,2,3,4]

// Assert function
assertEqual(len(a),4);

// Return hashcode
int hashcode = hash(key);
```


## MapString

`HashMap` where both key and value are string.

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


## Contact author


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 

