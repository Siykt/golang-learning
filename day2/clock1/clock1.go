package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 构造一个 listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// 循环监听连接
	for {
		// 接收连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// 处理连接
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// 关闭连接
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
