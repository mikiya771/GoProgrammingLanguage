package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, fmt.Sprint(start))
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Print(time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, dstName string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	nbytes, err := io.Copy(dst, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
