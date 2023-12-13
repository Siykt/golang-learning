package main

import "fmt"

// 包内变量包内可见
const boilingF = 212.0

func main() {
	// 函数内变量函数内可见
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
