package main

import (
	"fmt"
	"github.com/maxlcoder/gopl/ch2/tempconv"
)

func main() {
	var c tempconv.Celsius = 0
	var f tempconv.Fahrenheit = 0
	var k tempconv.Kelvin = 0

	fmt.Printf("%s = %s\n", c, tempconv.CToF(c)) // 这里触发隐式调用 String
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	fmt.Printf("%s = %s\n", k, tempconv.KToC(k))

}