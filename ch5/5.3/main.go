// 输出html文档树中的内容
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html_content: %v\n", err)
		os.Exit(1)
	}

	content(nil, doc)



}

func content(stack []string, n *html.Node)  {
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content(stack, c)
	}
}
