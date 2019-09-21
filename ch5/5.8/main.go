package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func forEachNode2(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) *html.Node {
	if pre != nil && pre(n, id) {
		return nil
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		resp := forEachNode2(c, pre, post, id)
		if resp == nil {
			return nil
		}
	}

	if post != nil && post(n, id) {
		return nil
	}
	return n
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
		depth++
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				fmt.Println("find")
				return true
			}
		}
	}
	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode2(doc, startElement, endElement, id)
}

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
	id := "toc"
	findResp := ElementByID(doc, id)
	fmt.Println(findResp)
}