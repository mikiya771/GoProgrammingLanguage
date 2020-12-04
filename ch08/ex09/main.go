package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walk(root, &n)
	}
	n.Wait()
}
func walk(dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	n.Add(1)
	go walkDir(dir, &n, fileSizes)

	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(dir, nfiles, nbytes)
		}
	}

	printDiskUsage(dir, nfiles, nbytes)
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("dir: %s  %d files %.1f GB\n ", dir, nfiles, float64(nbytes)/1e9)
}
