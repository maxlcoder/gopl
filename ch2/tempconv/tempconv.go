// 包 tempconv 进行摄氏温度和华氏温度的转换计算
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Inch float64
type Meter float64

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
