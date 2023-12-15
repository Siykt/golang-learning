package tempconv

func C2F(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }    // 摄氏转华氏
func F2C(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) } // 华氏转摄氏
