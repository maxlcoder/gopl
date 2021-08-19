package main

import (
	"fmt"
	"github.com/maxlcoder/gopl/ch8/8.10/links"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	cancelled := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancelled)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancelled)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string, cancelled chan struct{}) []string  {
	fmt.Println(url)
	list, err := links.Extract(url, cancelled)
	if err != nil {
		log.Print(err)
	}
	return list
}
