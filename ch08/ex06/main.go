package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

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
	flag.Parse()
	worklist := make(chan *depthList)
	var n int

	n++
	go func() { worklist <- &depthList{0, os.Args[1:]} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(list.depth, link)
			}
		}
	}
}
