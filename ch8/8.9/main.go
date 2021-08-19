package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// wakjDir 递归地遍历以 dir 为根目录地整个文件树
// 并在 fileSizes 上发送每个已找到地文件地大小

func walkDir(d string, n *sync.WaitGroup, root int, info chan<- dir)  {
	defer n.Done()
	for _, entry := range dirents(d) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(d, entry.Name())
			go walkDir(subdir, n, root, info)
		} else {
			info <- dir{root, entry.Size()}
		}
	}
}

// sema 是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回 dir 目录中地条目
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} // 获取令牌
	defer func() { <-sema }() // 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

type dir struct {
	id int
	size int64
}

func main() {
	// 确定初始目录
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	info := make(chan dir)
	var n sync.WaitGroup
	for id, root := range roots {
		n.Add(1)
		go walkDir(root, &n, id, info)
	}
	go func() {
		n.Wait()
		close(info)
	}()

	tick := time.Tick(500 * time.Millisecond)
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case dir, ok := <-info:
			if !ok {
				break loop // fileSizes 关闭
			}
			nfiles[dir.id]++
			nbytes[dir.id] += dir.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)

		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes[]int64)  {
	for id, root := range roots {
		fmt.Printf("%d files %.1f GB in %s\n", nfiles[id], float64((nbytes[id])/1e9), root)
	}
}
