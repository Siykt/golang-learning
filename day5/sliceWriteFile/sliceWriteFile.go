package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	writeBigFilepath := `C:\www\golang-learning\testData\writeBigFile.txt`

	// 312MB bigFile.txt
	bigFilepath := `C:\www\golang-learning\testData\bigFile.txt`

	bigFile, err := os.Open(bigFilepath)

	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	defer bigFile.Close()

	// 创建写入文件
	// flag os.O_CREATE|os.O_APPEND 为创建|追加
	// os.O_CREATE 如果文件不存在则创建
	// os.O_APPEND 如果文件存在则追加
	// perm 为权限参数, 0666 表示可读可写
	writeBigFile, err := os.OpenFile(writeBigFilepath, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	defer writeBigFile.Close()

	// 切片, 4096 为默认缓冲区大小
	buffer := make([]byte, 4096)

	for {
		// 读取文件
		n, err := bigFile.Read(buffer)

		// EOF 错误表示文件末尾
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("read file error: ", err)
			return
		}

		// 写入文件
		writeBigFile.Write(buffer[:n])
	}

	fmt.Println("write file success")
}
