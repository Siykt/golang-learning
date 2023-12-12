# Day 1

Go 语言的特点就是简洁, 从语法上直观的感受就是比较像 python.

## 特性

- Go 语言的代码通过包（package）组织, 需要声明 `package main` 来确定改文件为程序入口, Go 将会执行其 `func main` 方法.
- 可直接使用 `+` 链接字符串
- 一个文件夹下只能有一个入口也就是 `package main`

## 语法

### 变量与常量

变量声明方式:

```go
// 隐式初始化声明
var name string

// 显式初始化声明
var name string = "AntPro"

// 只能在函数中使用的快速声明
name := "AntPro"
```

### `for` 循环

Go 只有 `for` 循环, 它的语法为:

```go
for initialization; condition; post {
  // zero or more statements
}
```

其中 `initialization` 为**可选的** **初始化表达式**

`condition` 为**可选的** **循环条件**

`post` 为循环体执行结束后执行的语句, 在之后再次对 `condition` 求值

### map

语法:

map[key type]value type

实例:

```go
hash := make(map[int]int)
```

取值:

```go
// value 为值
// ok 为是否存在于map中
if value, ok := hash[key]; ok {
  // zero or more statements
}
```

赋值:

```go
hash[key] = value
```

### condition

必须是 boolean 值, 没有类型转换

- string 判断 `str == ""`
- char 判断 `str[i] == ' '`
