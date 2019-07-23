package main

import "fmt"

func main() {
	var a1 uint8 = 1<<1
	fmt.Println(a1)
	var a2 uint8 = 1<<5
	fmt.Println(a2)
	var a3 uint8 = 1<<1 | 1<<5

	fmt.Println(a3)

}
