package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)
var depthFlag = flag.Int("depth", math.MaxInt32, "depth of links")

type depthList struct {
	depth int
	list  []string
}

func crawl(depth int, url string) *depthList {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return &depthList{depth + 1, list}
}

func main() {
	Mirror(os.Args[1])
}

func saveURL(url, pathName string, isHTML bool) (linkMap map[string]string, err error) {
	fmt.Printf("start download %s as %s", url, pathName)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	f, err := os.Create(pathName)
	if err != nil {
		return
	}
	defer f.Close()
	if isHTML {
		changeMap, err := changeLink(resp)
		if err != nil {
			return nil, err
		}
	}
	for len(changeMap) != 0 {

	}
}

type urlType struct {
	url string
	tag string
}

func changeLink(resp *http.Response) (linkMaps map[string]urlType, err error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	nodes := []*html.Node{doc}
	for len(nodes) != 0 {
		for _, n := range nodes {
			if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
				for i, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					if strings.HasPrefix(link.String(), fmt.Sprintf("%s://%s", resp.Request.URL.Scheme, resp.Request.URL.Host)) {
						n.Attr[i].Val = "file:" + a.Val
						if a.Val != "/" && a.Val != "#" {
							linkMaps[a.Val] = urlType{link.String(), n.Data}
						}
					}
				}
			}
			if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img") {
				for i, a := range n.Attr {
					if a.Key != "src" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					n.Attr[i].Val = "file:" + a.Val
					linkMaps[a.Val] = urlType{link.String(), n.Data}
				}
			}
		}
	}
}
