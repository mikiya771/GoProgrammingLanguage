package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	getTextNode(doc)
}
func getTextNode(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.Data == "script" || n.Data == "style" {
			return
		}
	case html.TextNode:
		fmt.Println(strings.TrimSpace(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getTextNode(c)
	}
}
