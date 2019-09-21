package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(0, 0, doc)
	return
}

func countWordsAndImages(wordscount, imagescount int, n *html.Node) (words, images int)  {
	if n.Type == html.ElementNode && n.Data == "img" {
		imagescount += 1
	}
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordscount += 1
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordscount, imagescount = countWordsAndImages(wordscount, imagescount, c)

	}
	return wordscount, imagescount
}

func main() {
	url := "https://golang.org"
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Println("%s", err)
	}
	fmt.Println(words)
	fmt.Println(images)

}
