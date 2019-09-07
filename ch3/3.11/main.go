// 增加浮点数，正负数判断
// 思路是分割为三个部分，1.符号位，2整数位， 3，小数位 单独对整数位进行处理
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	start := 0
	// 正负号
	if s[0] == '+' || s[0] == '-' {
		start = 1
		buf.WriteByte(s[0])
	}
	// 分割
	end := strings.LastIndex(s, ".")
	fmt.Println(end)
	if end == -1 {
		end = len(s)
	}
	// 整数位
	intStr := s[start:end]
	n := len(intStr)
	for i, v := range intStr {
		// 起始位判断
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString("，")
		}
		buf.WriteRune(v)
	}
	// 写入小数点+小数位
	buf.WriteString(s[end:])
	return buf.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
