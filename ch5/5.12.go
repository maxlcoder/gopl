// forEachNode 调用pre(x) 和post(x) 遍历已n为根的书中的每个节点x
// 两个函数是可选的
// pre在子节点被访问前（前序）调用
// post在访问后（后序）调用
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
		}
	}
	var forEachNode func(n *html.Node, pre, post func(n *html.Node)) // 进行递归调用时，必需先声明，如果没有，那么递归调用的函数将不会是本身

	forEachNode = func(n *html.Node, pre, post func(n *html.Node)) {
		if pre != nil {
			pre(n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}

		if post != nil {
			post(n)
		}
	}
	forEachNode(doc, startElement, endElement)
}
