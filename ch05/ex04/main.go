package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range getFileLinks(nil, doc) {
		fmt.Println(link)
	}
}

func getFileLinks(str []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "img" || n.Data == "script" || n.Data == "style" {
			for _, a := range n.Attr {
				str = append(str, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str = getFileLinks(str, c)
	}
	return str
}
