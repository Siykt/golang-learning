package main

import "fmt"

func main() {
	var i int
	var b bool
	var s string
	var f float64

	// 简短变量声明
	a := 1

	fmt.Println(i, b, s, f, a)

	// 简短变量声明语句中必须至少要声明一个新的变量
	a1, b2 := fn()
	fmt.Println(a1, b2)
}

func fn() (int, int) {
	return 1, 2
}
