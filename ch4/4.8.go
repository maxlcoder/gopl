// charcount 计算 Unicode 字符的个数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode 字符数量
	var utflen [utf8.UTFMax + 1]int // UTF-8 编码长度
	invalid := 0                    // 非法 UTF-8 字符数量
	categoryCounts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	fmt.Println("统计开始")
	for {
		r, n, err := in.ReadRune() // 返回 rune, nbytes, error
		if err == io.EOF || r == '\n' { // 判断回车键退出循环
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// 统计字母，数字和其他三类即可，不用全部
		switch {
		case unicode.IsLetter(r):
			categoryCounts["letter"]++
		case unicode.IsDigit(r):
			categoryCounts["digit"]++
		default:
			categoryCounts["other"]++
		}
		counts[r]++
		utflen[n]++

	}
	fmt.Print("\nlen\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nlen\tcount\n")
	for c, n := range categoryCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
