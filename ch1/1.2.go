// 练习 1.2
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg) // println 默认在参数之间加入空格
		// or
		// fmt.Printf("index: %d, arg: %s", i, arg)
	}
}



