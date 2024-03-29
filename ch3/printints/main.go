// intsToString 与 fmt.Sprint(values) 类似，但插入逗号
package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString("，")
		}
		fmt.Fprintf(&buf, "%d", v)
		//buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "【1， 2， 3】"
}
