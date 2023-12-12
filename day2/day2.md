# Day 2

## Web 服务

`net/http` 库是 `go` 内置的 web 服务器库

通过调研 `http.ListenAndServe` 开启服务, 参数为: `http.ListenAndServe(addr string, handler http.Handler)`.
开启本地服务使用 `localhost:port` 即可.

构建路由处理器使用 `http.HandleFunc` 实现:

```go
// path, handler
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // 将URL地址写入response
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
})
```

## Go 的并发编程

### CSP 顺序通信进程

使用关键词 `chan` 创建数据管道:

```go
ch := make(chan string)
```

然后使用 `go` 关键词创建 `goroutine`:

```go
go fetch(url, ch)
```

当一个 `goroutine` 尝试在一个 `channel` 上做 send 或者 receive 操作时，这个 `goroutine` 会阻塞在调用处，直到另一个 `goroutine` 从这个 `channel` 里接收或者写入值，这样两个 `goroutine` 才会继续执行 `channel` 操作之后的逻辑
