package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	r io.Reader
	n int64
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if len(p) > int(lr.n) {
		p = p[:lr.n] // 读取lr.n个值
	}
	n, err = lr.r.Read(p)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader  {
	lr := LimitedReader{
		r: r,
		n: n,
	}
	return &lr
}

func main() {
	r := strings.NewReader("hello world! hello world!")
	lr := LimitReader(r, 8)
	buf := make([]byte, 10)
	n, _ := lr.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf))


}
