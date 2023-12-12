package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// /count 与 / 先后顺序无所谓?
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 并发时更新需要锁定资源, 处理竞态条件
		mu.Lock()
		count++
		mu.Unlock()

		fmt.Fprintf(w, "RES => %q\n", r.URL.Path)
	})

	http.ListenAndServe("localhost:8081", nil)
}
