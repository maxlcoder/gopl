package main

import (
	"fmt"
	"github.com/maxlcoder/gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	for _,arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		i := tempconv.Inch(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c)) // 温度转换
		fmt.Printf("%s = %s, %s = %s\n", i, tempconv.IToM(i), m, tempconv.MToI(m)) // 长度转换
	}
}