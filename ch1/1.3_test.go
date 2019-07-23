// 练习 1.3
package main

import (
	"os"
	"strings"
	"testing"
)

func RangeJoin() {
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
}

func StringJoin()  {
	strings.Join(os.Args[1:], " ")
}

func BenchmarkRangeJoin(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		RangeJoin()
	}
}

func BenchmarkStringJoin(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}

func main() {
	
}


// go test 1.3_test.go -bench=.
// goos: darwin
// goarch: amd64
// BenchmarkRangeJoin-4            100000000               12.6 ns/op
// BenchmarkStringJoin-4           500000000                3.74 ns/op
// PASS
// ok      command-line-arguments  3.523s
