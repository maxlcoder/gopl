// 去除相邻重复字符串
package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o", "o", "o", "f"}
	r := removeMulti(s)
	fmt.Println(r)
}

func removeMulti(s []string) []string {
	var r []string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			r = append(r, s[i])
		} else if s[i] != s[i+1] {
			r = append(r, s[i])
		}
	}
	return r
}
