// IntSet 是一个包含非负整数的集合
// 零值代表空的集合
package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has 方法的返回值表示是否存在非负数 x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加非负数 x 到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 将会对s和t做并集并将结果存在s中

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectionWith 交集
func (s *IntSet) IntersectionWith(t *IntSet)  {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith 差集
func (s *IntSet) DifferenceWith(t *IntSet)  {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Difference 差集
func (s *IntSet) Difference(t *IntSet)  {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
	}
}


// String方法以字符串"{1, 2, 3}"的形式返回集中

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)

			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	t := IntSet{}
	copy(t.words, s.words)
	return &t
}

func (s *IntSet) AddAll(x ...int) {
	for _, x := range x {
		s.Add(x)
	}
}

func main() {
	var x, y IntSet

	x.Add(2)
	x.Add(9)
	x.Add(10)
	x.Add(144)

	y.Add(1)
	y.Add(2)
	y.Add(9)



	x.Difference(&y)
	fmt.Println(x.String())

}
