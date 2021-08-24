package count

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

type count struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
	err     error
}

func charCount(r io.Reader) count {
	var c count

	c.counts = make(map[rune]int)
	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.err = err
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		c.counts[r]++
		c.utflen[n]++
	}
	return c
}
