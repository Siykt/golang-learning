package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://www.baidu.com"

	res, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Response Status ->", res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Response Body:\n%s", body)
}
