// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string
	fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
