package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%s", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Println(">")
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Println("/>")
		}
	}
}

func main() {
	for _, url := range os.Args[1:] {
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
		forEachNode(doc, startElement, endElement)
	}
}
