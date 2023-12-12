package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 创建一个通道
	ch := make(chan string)

	// 构造请求列表
	urls := []string{
		"https://www.baidu.com",
		// "https://www.google.com",
		"https://www.sina.com.cn",
		"https://www.qq.com",
		"https://www.163.com",
		"https://www.sohu.com",
		"https://www.ifeng.com",
	}

	for idx, url := range urls {
		// 并发发起请求, go关键字创建一个goroutine, 一个goroutine就是一个并发执行的函数
		fmt.Println("REQ", idx, "Start")
		go fetch(url, ch)
		fmt.Println("REQ", idx, "End")
	}

	// 输出结果
	for range urls {
		// ch通道接收到数据后, 打印结果, 通道接收数据是阻塞的, 所以这里会等待请求完成后才会输出结果
		// 如果不等待, 则会在请求发起后立即输出结果, 但是结果是乱序的
		fmt.Println("[RES]:", <-ch)
	}
}

func fetch(url string, ch chan<- string) {
	// 计算开始时间
	start := time.Now()

	// 发起请求
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// 读取响应内容
	body, err := io.Copy(ioutil.Discard, res.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	// 计算耗时
	secs := time.Since(start).Seconds()

	// 输出结果
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, body, url)
}
