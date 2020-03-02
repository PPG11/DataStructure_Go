# Go中方法的函数类型

今天在复现数据结构代码的时候，遇到了如下问题

```go
func (T Vector) traverse(fun1 func(int, int)string) {}
```

具体是在对Vector结构体，构建遍历方法(traverse)时，想要使用回调函数的办法，即通过将执行的函数传入`tranverse()`，进行执行函数。

但是这里遇见了一个问题，如果传入的函数本身是一个"方法(method)"，而不是一个"函数"(function)，那么这个"方法"是什么类型呢？

一开始模拟函数构造时的写法，简单的当成一个函数类型

```go
func (T Vector) traverse(fun1 (Vector)func(int, int)string) {}
```

发现编译不通过，出现了类型错误，意识到"方法"和"函数"是不一样的

那么对于结构体的方法到底是什么类型呢？尝试编写了如下代码

```go
package main

import (
	"fmt"
)

func main() {
	// 检查普通函数类型输出有没有问题
    var a func(int, int) string
	fmt.Printf("a: %T\n", a)
	// 检查将结构体作为第一个参数传入后是什么类型
	fmt.Printf("myF11: %T\n", myF11)
	// 检查对于结构体的方法又是什么类型
	fmt.Printf("vec.myF12: %T\n", vec.myF12)
	// 以上操作发现myF11和myF12是同样的类型
	
    // 所以想到是不是"方法"就是一种"函数"
    // 只是方法帮我们隐去了第一个变量
    // 做了如下尝试
	var b1, b2 func(vec, int) int
	b1 = myF11
	fmt.Printf("b1: %T\n", b1)
	b2 = vec.myF12
	fmt.Printf("b2: %T\n", b2)
}

type vec struct {
	int
	string
}

func myF11(t vec, a int) int {
	return t.int+a+1
}

func (t vec) myF12(a int) int {
	return t.int+a+1
}
```

发现编译通过，输出结果也和设想的一样
```go
a: func(int, int) string
myF11: func(main.vec, int) int
vec.myF12: func(main.vec, int) int
b1: func(main.vec, int) int
b2: func(main.vec, int) int
```

证明方法就是一种函数，只是替我们隐去了第一个变量，这里第一个变量默认为调用方法的结构体结构体本身

即一下两个函数完全是等价的()
```go
func myFun1(T myStruct, parm1 int, parm2 string) int
```

```go
func (T myStruct) myFun2(parm1 int, parm2 string) int
```