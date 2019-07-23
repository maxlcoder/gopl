// squares 函数返回一个函数，后者包含下一次要用到的平方数
// the nex square number each time it is called
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		fmt.Println(x)
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
