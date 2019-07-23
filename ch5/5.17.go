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
	findResp := ElementsByTagName(doc, "img")
	fmt.Println(len(findResp))
	for _, v := range findResp {
		fmt.Println(v.Attr)
	}
	// go run 5.17.go http://gopl.io
	// 总共4张图
	// [{ border 1} { src cover.png} { height 600}]
	// [{ border 0} { width 150} { src buyfromamazon.png}]
	// [{ border 0} { height 26} { src informit.png}]
	// [{ border 0} { width 150} { src barnesnoble.png}]




}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node  {

	startElement := func(n *html.Node, name string) (stop bool) {
		if n == nil {
			return false
		}
		if n.Type != html.ElementNode {
			return false
		}
		if n.Data == name {
			return true
		}
		return false
	}

	var nodes []*html.Node

	var forEachNode func(n *html.Node, name string, pre func(n *html.Node, name string) bool) // 进行递归调用时，必需先声明，如果没有，那么递归调用的函数将不会是本身

	forEachNode = func(n *html.Node, name string, pre func(n *html.Node, sname string) bool) {
		if pre != nil && pre(n, name) {
			nodes = append(nodes, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, name, pre)
		}

	}
	for _,name := range names {
		forEachNode(doc, name, startElement)
	}
	return nodes
	
}

