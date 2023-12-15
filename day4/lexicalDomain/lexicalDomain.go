package main

import (
	"log"
	"os"
)

func main() {
	hash := map[string]int{"a": 1}

	// 此时创建了一个隐式的词法域，v的作用域只在if语句块内
	if v, ok := hash["a"]; !ok {
		println("key a not exist")
	} else if v == 1 { // 此处的v是if语句块内的v
		println("v == 1")
	}
}

var cwd string

func bad() {
	// 此时, cwd将会被当成 init 内部的变量, 从而无法实现更新 cwd 的错误
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

// 解决方案为重新赋值或不使用 := 快捷声明
func g1() {
	// 暂存
	temp, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	// 重新赋值
	cwd = temp
}

func g2() {
	var err error
	// 不使用 :=, 直接赋值
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
