package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, tags ...string) (nodes []*html.Node) {
	if doc == nil {
		return
	}
	if doc.Type == html.ElementNode {
		for _, tag := range tags {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}
	nodes = append(nodes, ElementByTagName(doc.FirstChild, tags...)...)
	nodes = append(nodes, ElementByTagName(doc.NextSibling, tags...)...)
	return
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: findElement url id")
		os.Exit(1)
	}

	findElement(os.Args[1], os.Args[2])
}

func findElement(url string, tags ...string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	nodes := ElementByTagName(doc, tags...)
	for _, node := range nodes {
		fmt.Println(node.Data, node.Attr)
	}
}

func printNode(n *html.Node) {
	fmt.Printf("<%s", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s='%s'", a.Key, a.Val)
	}
	fmt.Println(">")
}
