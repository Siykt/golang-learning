package main

import (
	"fmt"

	"antpro.me/day4/tempconv"
)

func main() {
	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	fmt.Println(c == 0)                   // true
	fmt.Println(f >= 0)                   // true
	fmt.Println(c == tempconv.Celsius(f)) // true
	// fmt.Println(c == f) // 编译错误：类型不匹配

	fmt.Println(tempconv.C2F(tempconv.BoilingC))      // 212
	fmt.Println(tempconv.C2F(tempconv.FreezingC))     // 32
	fmt.Println(tempconv.C2F(tempconv.AbsoluteZeroC)) // -459.66999999999996
	fmt.Println(tempconv.F2C(212.0))                  // 100°C
	fmt.Println(tempconv.F2C(32.0))                   // 0°C
	fmt.Println(tempconv.F2C(-459.67))                // -273.14999999999986°C
}
