# Day 3

## 声明

go 的声明方式分为 4 种:

1. var 变量声明
2. const 常量声明
3. type 类型声明
4. func 函数声明

在单个 go 源文件中, 直接在文件内声明的值本文件中都可访问:

```go
package main

// 包内均可访问
var a = 1
```

在函数体内声明的值仅可在该函数中访问:

```go
package main

func a() {
  var b = 2
  return b
}

func main() {
  b // error
}
```

### 变量

go 声明变量的语法如下

```go
var variableName type = value
```

其中 `type` 或赋值的部分可以省略其一:

```go
// 忽略显式类型 int
var a = 1

// 默认赋值 0
var b int
// 默认赋值 ""
var s string
// 默认赋值 false
var bool bool
```

还有一种只能在函数内使用的简短变量声明, 就是省略 `var` 使用 `:=` 操作符:

```go
func fn() {
  a, b := 1, 2
  fmt.Println(a, b)
}
```

简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：

```go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables
```

### 指针

go 的指针没有 c 那么复杂, 主要的功能就是提供**引用内存地址**的功能, 可以类比为 js 的对象.

在声明时使用 `&` 符号即可创建指针, 其指向地址为 `&` 跟随变量的地址:

```go
func main() {
  x := 1

  // 引用 x 的内存地址
  p := &x
}
```

对指针使用 `*` 符号可以从指针中获取实际变量的值:

```go
fmt.Println(*p) // 等于 x 的值
```

同时也能使用 `*` 符号修改实际变量的值:

```go
*p = 2
fmt.Println(x) // 等于 *p 也就是 2
```

那么函数就可以通过传入指针的方式修改传入值:

```go
func incr(p *int) int {
  *p++
  return *p
}

var x = 1

incr(&x) // x = 2
```

### new 函数

使用 new 函数可以创建该类型的指针, 语法为 `new(T)`, 实例如下:

```go
p := new(int)
fmt.Println(*p) // 0 默认值

*p = 2
fmt.Println(*p) // 2
```
