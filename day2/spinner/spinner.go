package main

import (
	"fmt"
	"time"
)

func main() {
	// 开启一个goroutine, 执行 spinner 函数
	go spinner(100 * time.Millisecond)

	// 计算斐波那契数列第46个数的值
	const n = 46
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	// 循环打印 spinner 字符
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// var cache = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		// cache[n] = n
		return n
	}
	return fib(n-1) + fib(n-2)
	// if _, ok := cache[n]; !ok {
	// 	cache[n] = fib(n-1) + fib(n-2)
	// }
	// return cache[n]
}
