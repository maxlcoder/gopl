// 统计sha256散裂中不同位的个数
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(bytecount([]byte{1, 3, 5, 8}, []byte{1, 2, 6}))
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(bytecount(c1[:], c2[:]))
}

func bytecount(b1, b2 []byte) int {
	len1, len2 := len(b1), len(b2)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}
	count := 0
	for i := 0; i < maxLen; i++ {
		if i >= len1 || i >= len2 {
			count += 8 // 超出的部分，默认每个 byte = 8bit
		} else {
			count += bitcount(b1[i], b2[i])
		}
	}
	return count
}

func bitcount(b1, b2 byte) int {
	// 异或，计算1的个数
	temp := b1 ^ b2
	count := 0
	for temp != 0 {
		temp &= temp - 1 // 清除最低位的1
		count++
	}
	return count
}
