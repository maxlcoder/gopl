package tempconv

// CToF 把摄氏温度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 把华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func IToM(i Inch) Meter {
	return Meter((i * 2.54) / 100)
}

func MToI(m Meter) Inch {
	return Inch(m * 100 / 2.54)
}