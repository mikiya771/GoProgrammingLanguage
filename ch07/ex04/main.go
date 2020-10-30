package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func Parse(c string) {
	doc, err := html.Parse(newReader(c))
	if err != nil {
		return
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
func newReader(c string) io.Reader {
	return &reader{[]byte(c), 0}
}

type reader struct {
	bytes []byte
	next  int
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if r.next >= len(r.bytes) {
		return 0, io.EOF
	}
	nBytes := len(r.bytes) - r.next
	if nBytes > len(p) {
		nBytes = len(p)
	}

	copy(p, r.bytes[r.next:r.next+nBytes])
	r.next += nBytes
	return nBytes, nil
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
