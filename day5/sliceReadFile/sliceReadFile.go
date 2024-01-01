package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := `C:\www\golang-learning\testData\bigFile.txt`
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// scanner.Text() 返回当前行的内容
		fmt.Println(scanner.Text())
		// 每次回车输出下一行
		bufio.NewReader(os.Stdin).ReadRune()
	}
}
