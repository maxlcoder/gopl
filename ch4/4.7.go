package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("I said 北京您早")
	var i []string
	for _, v := range string(s) {
		i = append(i, string(v))
	}
	reverseString(i) // 转换了类型，同时空间内存空间也重新分配了，不符合标准
	fmt.Println(i)

	reverseUTF8(s)
	fmt.Printf("%q", s)
}

func reverseString(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		if size > 1 {
			reverse(b[i : i+size]) // 局部字符反转
		}
		i += size
	}
	reverse(b)
	return b
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
