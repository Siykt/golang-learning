package main

import (
	"fmt"
	"net/http"
)

func main() {
	// path, handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 将URL地址写入response
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	http.ListenAndServe("127.0.0.1:8081", nil)
}
