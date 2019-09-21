package main

import "fmt"

func unZero() (x int) {
	defer func() {
		if err := recover(); err != nil {
			x = 1
		}
	}()
	panic("测试")
}

func main() {
	resp := unZero()
	fmt.Println(resp)
}