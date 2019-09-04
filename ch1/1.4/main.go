// dup2 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

type counts map[string]int // 定义类型

func main() {
	fileCount := make(map[string]counts)
	files := os.Args[1:] // 文件
	if len(files) == 0 {
		countLines(os.Stdin, fileCount)
	} else {
		for _, arg := range files { // 记录file
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCount)
			f.Close()
		}
	}
	for file, count := range fileCount {
		for line, n := range count {
			if n > 1 {
				fmt.Println(file)
				fmt.Printf("%d\t%s\n", n, line)
				break // 只要只要有一个多行，继续下一个文件
			}
		}
	}
}

func countLines(f *os.File, fileCount map[string]counts) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	fileCount[f.Name()] = counts
	// 注意：忽略 input.Err() 中可能的错误
}
