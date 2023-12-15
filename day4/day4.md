# Day4

## Golang 中变量的生命周期

在源文件声明的变量与整个程序的运行周期是一致:

```go
package main

// 此变量将在程序运行结束后被回收
var str = "str"
```

所有变量的声明方式都有函数作用域与块级作用域, 正常情况下函数或脱离块就不可访:

```go
for i := 0; i < 1; i++ {
  // 变量 i 仅在 for 循环内部可见
}

// i
// 外部访问时go编译器提示 undeclared name: i
```

当变量被外部指针引用后将超越函数作用域与块级作用域的限制:

```go
package main

var global *int

func main() {
	for i := 0; i < 10; i++ {
		// 变量 i 仅在 for 循环内部可见

		// 使用指针脱离变量作用域限制, 使i于外部可访
		global = &i

		println("setGlobal Before ->", i) // 0, 2, 4, 6, 8

		incrGlobal()

		println("setGlobal After ->", i) // 1, 3, 5, 7, 9
	}

	// i
	// 外部直接访问时go编译器提示 undeclared name: i
}

func incrGlobal() {
	// 越权访问 i
	*global += 1
}

```

Golang 的垃圾回收机制与 V8 类似, 都是以可达性(Mark-Sweep GC)来判断是否可被回收的.

扩展阅读:

1. <https://go.dev/blog/ismmkeynote>
2. <https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/#72-%E5%9E%83%E5%9C%BE%E6%94%B6%E9%9B%86%E5%99%A8>

## 赋值

对于两个值是否可以用 `==` 或 `!=` 进行相等比较的能力也和可赋值能力有关系:
对于任何类型的值的相等比较, 第二个值必须是对第一个值类型对应的变量是可赋值的, 反之亦然

自增和自减是语句, 不是表达式, 所以不能支持如下语法:

```go
x = i++
```

### 元组赋值

Golang 支持元组赋值的便捷方式以便以更紧凑的赋值变量:

```go
// 斐波那契数列
func fib(n int) int {
	// 动态规划 + 空间压缩
	x, y := 0, 1

	for i := 0; i < n; i++ {
		// 交换计算结果
		x, y = y, x+y
	}

	return x
}
```

常见的 tuple 返回有几种:

1. 错误返回, 如 http.Get/os.Open 等: `res, err := http.Get(url)`
2. 条件返回, 如 map 查询/类型断言/通道数据接收等
   1. `value, ok := map[key]` map 查询
   2. `value, ok := x.(T)` 类型断言
   3. `value, ok := <- ch` 通道数据接收

## 类型

类型是为了分隔不同概念而定义一个新的/独立的类型, 这使得两个底层类型相同的类型无法兼容 (非鸭子类型)

语法: `type name 底层类型`, 实例:

```go
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
```

虽然 `Celsius` 与 `Fahrenheit` 具有相同的底层类型 `float64`,
但它们无法进行比较或混在一个表达式运算中, 不过它们可以用底层的`float64`类型进行比较与计算:

```go
var c Celsius
var f Fahrenheit

fmt.Println(c == 0) // true
fmt.Println(f >= 0) // true
fmt.Println(c == f) // compile error: type mismatch
// 类型转换
fmt.Println(c == Celsius(f)) // true
```

类型的转换方法实现须要它们的底层类型相同:

```go
func c2f(c Celsius) Fahrenheit {
  return Fahrenheit(c*9/5 + 32)
}
```

命名类型还可以为该类型的值定义新的行为, 比如最基本的 String() 功能:

```go
func (c Celsius) String() string {
  return fmt.Sprintf("%g°C", c)
}
```

## 包与源文件

创建一个包文件需要在根目录下创建 `go.mod` 声明文件:

```
//  定义包名
module antpro.me

// 定义 golang 版本
go 1.19
```

之后你即可通过 `package name` 的方式定义包, 如文件 `day4/tempconv/conv.go`:

```go
package tempconv

func C2F(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }    // 摄氏转华氏
func F2C(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) } // 华氏转摄氏
```

之后在别处使用即可通过导入 `antpro.me/day4/tempconv` 的方式获取包内导出的方法与变量,
而包内以大写开头声明的方法与变量即为导出 (没有显式声明 `export` 的方式)

### init 初始化方法

如果这个包需要复杂的初始化逻辑, 可以通过在源文件中设置 `init` 方法的方式实现:

```go
package popcount

// pc[i] is the population count of i.
var pc [256]byte

// 包的初始化逻辑
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

```

### 包的发布与托管

包可以托管在任何一个具有 Go 模块代理的服务器上如
[Athens](https://github.com/gomods/athens)/
[goproxy](https://github.com/goproxy/goproxy.cn)
, Github 提供了 Go 模块代理的功能可以直接通过它来托管 golang 包

具体步骤如下:

1. 修改 go.mod 为 `github.com/username/repo`, 如: `github.com/siykt/go-test-mod`
2. 上传代码至该地址
3. 发布一个 release 如 `v1.0.0`
4. 在其他地方使用 go get 命令下载包至项目, 如: `go get github.com/siykt/go-test-mod@latest`
5. 通过 `import` 导入包, 如: `import "github.com/siykt/go-test-mod"`

### VSCode 集成

VSCode 可以集成 `goimport` 实现保存时格式化, 并自动导入使用的包与删除没有使用的包,
集成方式如下:

#### 1. 安装 `goimports`

通过命令 go 安装:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

#### 2. 配置 `goipmorts`

设置 `goimports` 为默认的格式化工具, 在 vscode 的 settings 配置:

```json
"[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": true
    },
    "editor.defaultFormatter": "golang.go",
    "go.formatTool": "goimports"
}
```

## 作用域

Golang 的作用域与 js 不同, 主要体现在块级作用域上,
Golang 会提供一些隐式的作用域:

```go
// 此时创建了一个
func main() {
	hash := map[string]int{"a": 1}

	// 此时创建了一个隐式的词法域，v的作用域只在if语句块内
	if v, ok := hash["a"]; !ok {
		println("key a not exist")
	} else if v == 1 { // 此处的v是if语句块内的v
		println("v == 1")
	}
}
```

`:=` 会重新声明变量, 这将导致作用域链发生变化:

```go
var cwd string

func init() {
	// 此时, cwd将会被当成 init 内部的变量, 从而无法实现更新 cwd 的错误
	cwd, err := os.Getwd()
	if err != nil {
			log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
```

可以通过别名设置或不使用 `:=` 快速声明的方式:

> 修改后

```go
var cwd string

func init() {
	// 此时, cwd将会被当成 init 内部的变量, 从而无法实现更新 cwd 的错误
	cwd, err := os.Getwd()
	if err != nil {
			log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
```
