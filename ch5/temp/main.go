package main

import (
	"fmt"
	"os"
)

func main() {

	var rmdirs []func()
	for _, dir := range tempDirs() {
		dir := dir // 这里重新申明，可以保持一次循环是一个dir的副本，不会随着dir的更新，而更新副本
		fmt.Println(dir)
		os.MkdirAll(dir, 0775)
		rmdirs = append(rmdirs, func() {
			fmt.Println(dir)
			os.RemoveAll(dir)
		})
	}
	for _, rmdir := range rmdirs {
		rmdir();
	}
}

func tempDirs() []string {
	temp := make([]string, 0)
	temp = append(temp, "f1")
	temp = append(temp, "f2")
	return temp
}
