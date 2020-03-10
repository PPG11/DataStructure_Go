# `import` 其他包中结构体和方法

文件目录如下
```shell script
.
|-- Vector
|    |- Vector_base.go
|
|-- Stack
     |- Stack_base.go
```

需要在`Stack_base.go`中引用`Vector_base.go`中的结构体`Vector`，具体如下

在`Vector_base.go`中定义了相应结构体和方法

```go
package Vector

type Rank = int

type Vector struct {
	_size     Rank
	_capacity int
	_elem     []interface{}
}

func (T Vector) size() Rank {
	return T._size
}

func (T Vector) Empty() bool {
	if T._size == 0 {
		return true
	} else {
		return false
	}
}
```

注意到`size()`和`Empty()`都是结构体`Vector`的方法，
但是`size()`是小写字母开头，`Empty()`是大写字母开头

在`Stack_base.go`中`import`了该包

```go
import "datastructure/DataStructure_Go/Cpt_2_Vector/Vector"

// 注意这里
type Stack Vector.Vector

func find() {
    var T Stack
}
```

注意这里用的是

```go
type Stack Vector.Vector
```

即`Stack`是一个和`Vector`相同定义的新类型，仅仅是有相同定义而已
所以自然没有得到导入的包中所有相应的方法

如果改为
```go
type Stack = Vector.Vector
```

则 `Stack` 实为 `Vector` 的别名，即就是`Vector`，
所以可以按照原本的大小写规则访问到 `import` 包中能访问到的方法

本例中，首字母大写的 `Empty()` 方法可以被 `Stack` 使用

而首字母小写的 `size()` 方法不能使用

---

但是还有新问题！

### 在`import`结构体后，需要定义新的方法

会发现提示`cannot define new methods on non-local type`

非本地包无法添加新的方法，但是这里如果要添加 `Stack` 的一些新功能怎么办呢

这里采用的是"继承"的方法，即通过与如下设置得到

```go
type Stack struct {
    Vector.Vector
}
```

这样这里`Stack`可以通过 `Stack.Vector.xxx` 访问原来的函数

并且可以通过给Stack设置新的 `method` 来增加新功能

这个时候再思考这个过程，其实本身 `Stack` 也算是对 `Vector` 的一种继承

所以这样设置是合理的

bingo~