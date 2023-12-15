package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func c2f(c Celsius) Fahrenheit      { return Fahrenheit(c*9/5 + 32) }    // 摄氏转华氏
func f2c(f Fahrenheit) Celsius      { return Celsius((f - 32) * 5 / 9) } // 华氏转摄氏
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }    // 摄氏温度字符串
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }    // 华氏温度字符串

func main() {
	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // 编译错误：类型不匹配
	fmt.Println(c == Celsius(f)) // true

	fmt.Println(c2f(BoilingC))      // 212
	fmt.Println(c2f(FreezingC))     // 32
	fmt.Println(c2f(AbsoluteZeroC)) // -459.66999999999996
	fmt.Println(f2c(212.0))         // 100°C
	fmt.Println(f2c(32.0))          // 0°C
	fmt.Println(f2c(-459.67))       // -273.14999999999986°C
}
