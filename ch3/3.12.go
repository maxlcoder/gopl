// 判断同文异构
// 思路是：循环去除str1中的字符在str2中查找，找到则剔除，最后判断str2是否为空
package main

import (
	"fmt"
	"os"
	"strings"
)

func comp(x, y string) bool {
	if len(x) != len(y)  {
		return false
	}

	for _, s := range x {
		position := strings.LastIndex(y, string(s))
		fmt.Println(position)
		if position != -1 {
			y = y[:position] + y[position+1:] // 剔除元素
		}
		fmt.Println(y)
	}
	// 通过剩余长度判断
	if len(y) > 0{
		return false
	}
	return true

}

func main() {
	fmt.Println(comp(os.Args[1], os.Args[2]))
}
