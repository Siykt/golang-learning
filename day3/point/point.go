package main

import "fmt"

func main() {
	x := 1

	// &x 获取 x 的内存地址
	p := &x

	// *p 获取 p 指向的值, 即 x 的值
	fmt.Println(*p)

	// *p = 2 修改 p 指向的值, 即 x 的值
	*p = 2
	// 此时 x = 2
	fmt.Println(x)

	// p = 3 不能修改 p 的值, 因为 p 是一个指针变量, 不能直接赋值
	// p = 3 cannot use 3 (untyped int constant) as *int value in assignment

	// 使用 incr 同样可以修改变量x的值
	incr(p)
	incr(&x)
	fmt.Println(x)
}

func incr(p *int) {
	*p++
}
