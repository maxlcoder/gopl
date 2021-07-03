package tempconv

type Celsius float64 // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BollingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 +32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5 /9)
}