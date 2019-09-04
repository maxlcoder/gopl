package main

import (
	"fmt"
	"github.com/maxlcoder/gopl/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string) // 可能有重复的 URL 列表
	unseenLinks := make(chan string) // 去重后的 URL 列表

	// 向任务列表中添加命令行参数
	go func() { worklist <- os.Args[1:]}()
	// 创建 20 个爬虫 goroutine 来获取每个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	// 并发爬取 web
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
