// 包 tempconv 进行摄氏温度和华氏温度的转换计算
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Inch float64
type Meter float64

// *celsiusFlag 满足 flag.Value 接口
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit) // 无须检查错误
	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
		return nil
	case "F", "ºF":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "ºK":
		f.Celsius = KToC(Kelvin(value))
		return nil


	}
	return fmt.Errorf("invalid temperature %q", s)

}


// CelsiusFlag 根据给定的 name,默认值和使用方法
// 定义了一个Celsius标志，返回了标志值的指针
// 标志必须包含一个数值和一个单位，比如： "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gºK", k)
}

func (i Inch) String() string {
	return fmt.Sprintf("%g inch", i)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g meter", m)
}
