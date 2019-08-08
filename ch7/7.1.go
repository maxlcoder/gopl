package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	*c = 0
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	*c = 0
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	return int(*c), nil
}


func main() {
	var c WordCounter
	c.Write([]byte("hello world\n"))
	fmt.Println(c)

	var lc LineCounter

	lc.Write([]byte("hello world\n hello world\n hello world\n"))
	fmt.Println(lc)
}