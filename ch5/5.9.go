package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
   sArr := strings.Split(s,"$foo")
   resp := ""
	if len(sArr) > 1 {
		for k, v := range sArr {
			if k == len(sArr) - 1 {
				resp += v
			} else {
				resp += v + f("foo")
			}
		}
	} else {
		resp += sArr[0]
	}
   return resp
}

func f(s string) string {
	if s == "foo" {
		return "find foo"
	}
	return "find foo failed"
}

func main() {
	resp := expand("hello$fooand", f)
	fmt.Println(resp)
}


