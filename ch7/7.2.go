package main

import (
	"fmt"
	"io"
	"os"
)

type BytingCounter struct {
	w     io.Writer
	count int64
}

func (c *BytingCounter) Write(p []byte) (n int, err error) {
	c.w.Write(p)
	n = len(p)
	c.count += int64(n)
	return
}

func CountWriter(w io.Writer) (io.Writer, *int64) {
	c := BytingCounter{w, 0}
	return &c, &c.count
}

func main() {
	w, n := CountWriter(os.Stdout)
	fmt.Fprint(w, "hello world")
	fmt.Println(n)
}
