
```
 _     _   
 ___ ___ ___ ___|_|___| |_ 
| . |_ -|  _|  _| | . |  _|
|_  |___|___|_| |_|  _|_|  
|___|             |_|   

```

<div align="center">  

📘[特性](#特性) 🌰[例子](#例子)💡 [联系作者](#联系作者)|🇦🇺[英文文档](https://github.com/crossoverjie/gscript/blob/master/README.md)


</div><br>


# 介绍

这是一门用 Go 编写的一款**静态、强类型**的脚本语言，大部分语法参考了 Java 以及少量的 Go。

> 当前版本仅供学习与实验。


# 特性
- [x] class声明。
- [x] 函数声明与调用。
- [x] 基本类型: `int/string/float/bool`
- [x] 特殊类型 `nil`
- [x] 闭包：函数一等公民。
- [x] 内置函数: `len()/hash()/assertEqual()`
- [x] 标准库：`Map/LinkedList/Array`

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

更多样例请参考：[https://github.com/crossoverJie/gscript/tree/main/example](https://github.com/crossoverJie/gscript/tree/main/example)

## 联系作者


> crossoverJie#gmail.com

![qrcode_for_gh_3a954a025f10_258.jpg](https://i.loli.net/2019/07/09/5d245f3e955ce61699.jpg) 