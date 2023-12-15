package main

func main() {
	println(fib(6)) // 8
}

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
