// 就地反转一个整型 slice 中的元素
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 3, 3, 5}
	reverse(a[:])
	fmt.Println(a)
}
