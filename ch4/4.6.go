// 相邻的多个空格转换为一个空格
package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("h  ll e   o")
	fmt.Printf("%q", cutSpace(b))

}

func cutSpace(b []byte) []byte {
	var r []byte
	for i, v := range b {
		if i == 0 {
			r = append(r, v)
		} else if !unicode.IsSpace(rune(v)) {
			r = append(r, v)
		} else if unicode.IsSpace(rune(v)) && !unicode.IsSpace(rune(r[len(r)-1])) {
			r = append(r, ' ')
		}
	}
	return r
}
