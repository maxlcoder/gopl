package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// 读取输入的文件（多个）
	for _, filename := range os.Args[1:] {
		// 读取文件全部内容
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// 换行符分割文件，
		for _, line := range strings.Split(string(data), "\n") {
			// 每一行的内容作为key
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
