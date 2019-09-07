// 统计单词次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // Unicode 字符数量
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	fmt.Println("统计开始")
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Print("\nlen\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
