package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "省略行尾换行符")
var sep = flag.String("s", " ", "分隔符")

func main() {
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
