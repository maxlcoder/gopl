package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type StringReader struct {
	s string
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	n = len(p)
	err = io.EOF
	copy(p, []byte(sr.s))
	return
}

func NewReader(s string) *StringReader {
	return &StringReader{s: s}
}

func main()  {
	doc, _ := html.Parse(NewReader(`<html><head><title>xx</title></head><body><div>xxxdic</div></body></html>`))
	fmt.Println(doc.FirstChild.Data)
}
