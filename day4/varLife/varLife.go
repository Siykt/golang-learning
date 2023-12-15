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
