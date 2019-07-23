package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileNames := make(map[string]map[string]int) // 定义多维map，这里要注意多维map的使用
	// counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countFileLines(os.Stdin, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countFileLines(f, fileNames)
			f.Close()
		}
	}

	for file, count := range fileNames {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
				fmt.Println(file)
			}
		}



	}
}


func countFileLines(f *os.File, fileNames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	fileNames[f.Name()] = counts
}