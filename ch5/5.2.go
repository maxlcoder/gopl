package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	stack := make(map[string]int)
	stack = elementCount(stack, doc)
	fmt.Println(stack)

}

func elementCount(stack map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			stack["a"]++
		case "p":
			stack["p"]++
		case "span":
			stack["span"]++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stack = elementCount(stack, c)
	}
	return stack
}
