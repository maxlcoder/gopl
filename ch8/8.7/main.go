package main

import (
	"flag"
	"github.com/maxlcoder/gopl/ch5/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	base = flag.String("base", "http://www.baidu.com", "")
	wg   sync.WaitGroup
)

func main() {
	flag.Parse()
	for _, url := range crawl(*base) {
		wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			download(*base, url)
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	<-done
}

func download(base, url string) {
	if !strings.HasPrefix(url, base) {
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "https://")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}
	filename := dir + "index.html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
