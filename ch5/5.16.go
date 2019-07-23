package main

import "fmt"

func join(step string, vals ...string) string {
	temp := ""
	for i, val := range vals {
		if i == 0 {
			temp = val
		} else {
			temp = temp + step + val
		}

	}

	return temp
}

func main() {
	fmt.Println(join(",", "你好", "赵钱孙李"))
}