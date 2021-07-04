// 重写comma
package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, v := range s {
		// 起始位判断
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString("，")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
