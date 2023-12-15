package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	for _, arg := range os.Args[1:] {
		runNougat(arg)
	}
}

func runNougat(path string) {
	// cmd := exec.Command("conda", "activity", "stable-diffusion-webui")

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	cmd := exec.Command("C:\\Users\\siykt\\.conda\\envs\\stable-diffusion-webui\\Scripts\\nougat.exe", path, "-o", "./")

	// 创建管道
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 创建一个新的扫描器，用于读取标准输出
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 打印每一行输出
	}

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
