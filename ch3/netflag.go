package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota // 向上
	FlagBroadcast // 支持广播访问
	FlagLoopback // 是环回接口
	FlagPointToPoint // 属于点对点链路
	FlagMulticast // 支持多路广播访问
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	fmt.Printf("%b %t\n", v, IsCast(v))
}
