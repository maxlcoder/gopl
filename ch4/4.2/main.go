package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	sha := flag.Int("sha", 256, "sha hash lenght 256/384/513")
	s := flag.String("s", "", "string to sha hash")
	flag.Parse()
	var sha_hash []byte
	switch *sha {
	case 384:
		temp := sha512.Sum384([]byte(*s))
		sha_hash = temp[:]
	case 512:
		temp := sha512.Sum512([]byte(*s))
		sha_hash = temp[:]
	default:
		temp := sha256.Sum256([]byte(*s))
		sha_hash = temp[:]
	}
	fmt.Printf("sha%d(%q) 转换为 %x\n", *sha, *s, sha_hash)
}
