package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// fmt.Println(ZB) constant 1180591620717411303424 overflows int
	// fmt.Println(YB) constant 1208925819614629174706176 overflows int
}
