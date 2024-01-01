# Day 5

## Golang 的文件模块

Golang 通过 `os` 模块提供文件操作能力, 它的功能十分强大.

### 创建文件流

创建文件流可以使用 `os.Create("filename")` 方法实现,
它的特点是覆盖模式, 也就是说如果文件已经存在, go 将清空其内容.

实例如下:

```go
package main;

import (
  "os"
)

func main() {
  f, err := os.Create("str.txt")

  if err != nil {
    return
  }

  defer f.Close()
}
```

`defer` 语法保证了 main 函数退出时资源被释放.

当 `err != nil` 时, `f` 不会被创建, 所以我们也不需要也不应该去释放资源.

在此之后即可通过使用 `f.WriteString("String")` 方法来写入数据,
需要注意的是, 多次调用 `f.WriteString` 只会将字符串附加至文件底部, 而不是覆盖.

### 读取文件

我们可以使用 `os.Open("filepath")` 方法来读取文件:

```go
package main;

import (
  "os"
)

func main() {
  f, err := os.Open("str.txt")

  if err != nil {
    return
  }

  defer f.Close()
}
```

接下来我们同样需要使用 `f.Read` 方法来读取数据, 这时数据将会被保存在内存中(变量).
所有需要注意当文件数据非常大的时候可能会影响应用性能.
解决方案可以通过数据分片来避免.

### 大文件读写与数据切片

所有的大文件处理都是使用数据切片的技术, golang 也一样.

数据切片的核心逻辑是通过设置一个缓存区来避免一次性读取大量的数据至内存当中,
特别是我们都不需要使用到的数据.

golang 中可以使用 `make([]byte, size)` 来创建缓存区:

```go
buffer := make([]byte, 1024)
```

读取时将 buffer 传入 `f.Read`, 当切片长度小于内容时将直接读取所有的数据.

其中缓冲区的大小可以使用 **1024 byte(1 字节)** 或 **4096 byte(4 字节)**,
因为它们在不同的系统和应用中表现通常都不错.

而需要读取的数据内容长度大于切片的长度将仅返回切片长度的内容,
并且在下一次读取数据时使用上一次读取时的位置继续读取.

```go
f, err := os.Open("big_file_filepath")

buffer := make([]byte, 1024)

for {
  len, err := f.Read(buffer)
  if err != nil {
    break
  }

  fmt.Println(string(buffer[:len]))
}
```

值得注意的是, 当 `f.Read` 将所有的数据读取完成后将会抛出 `EOF` 的错误,
此时可以使用 `err == io.EOF` 来退出数据读取:

```go
len, err := f.Read(buffer)

if err == io.EOF {
  break
} else if err != nil {
  fmt.Println("read file error: ", err)
  return
}
```

除了 `buffer` 切片的方式外, 还有一种比较 golang 风格的方式,
是使用 `bufio.NewScanner` 来实现每行的扫描,
当然这种方式并不能处理一行多数据的情况:

```go
scanner := bufio.NewScanner(f)

for scanner.Scan() {
  // scanner.Text() 返回当前行的内容
  fmt.Println(scanner.Text())
  // 每次回车输出下一行(限制输出速率)
  bufio.NewReader(os.Stdin).ReadRune()
}
```

同理, 写入数据也是遵循以上的切片规则, 无非是将读取数据的逻辑切换成写入逻辑.
详细参考 `day5/sliceWriteFile/sliceWriteFile.go` 实现.
