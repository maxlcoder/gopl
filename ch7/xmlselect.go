// Xmlselect 输出 XML 文档中指定元素下的文本
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)

	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // 入栈
		case xml.EndElement:
			stack = stack[:len(stack)-1] // 出栈
		case xml.CharData:
			if containsAl {

			}
		}

	}
}

// containsAll 判断 x 是否包含 y 中的所有元素，且顺序一致
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
