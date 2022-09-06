


<div align="center">  

```
 _     _   
 ___ ___ ___ ___|_|___| |_ 
| . |_ -|  _|  _| | . |  _|
|_  |___|___|_| |_|  _|_|  
|___|             |_|   

```

📘[Features](#features) 🌰[Demo](#demo) 🎉[Syntax](#syntax) 🎁[Standard library](#标准库) 🔧[Install](https://github.com/crossoverJie/gscript/releases)💡 [Contact Author](#contact-author)|🇨🇳[中文文档](https://github.com/crossoverjie/gscript/blob/master/README-ZH.md)



</div><br>

# Introduction

This is a **statically and strongly** typed language written in Go, the syntax of Java and Go is referenced.

> The current version is for study and experimentation only.

```shell
❯ gscript example/hello_world.gs
hello world
```

# Features

- [x] Class declaration.
- [x] Function declaration and call.
- [x] Primitive type: `int/string/float/bool`.
- [x] Array type.
- [x] `nil` type.
- [x] Function type.
- [x] Closure：Functions as First-Class Objects.
- [x] Native function: `len()/hash()/assertEqual()`.
- [x] Standard library：`Map/LinkedList/Array`.
- [ ] Native support `json`.
- [ ] Native support `http`.



# Demo

```java
int b= 10;
int foo(int age){
	for(int i=0;i<10;i++){
		age++;
	}
	return b+age;
}
int add(int a,int b) {
	int e = foo(10);
	e = e+10;
	return a+b+3+e;
}
add(2,20);
// Output:65
```

# Syntax

# Standard lib

# Installation

```shell
go install github.com/crossoverJie/gscript/cmd@latest

gscript -h
```

Or download the [binary file](https://github.com/crossoverJie/gscript/releases) directly

# Features

- [x] Basic quadratic operations
- [x] Scope and function.
- [ ] More basic type
- [ ] Support scope and function
- [ ] Support closure
- [ ] OOP

![](doc/run.jpg)


## Contact author


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 

