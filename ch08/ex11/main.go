package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch)
	}
	fmt.Print(<-ch)
	close(done)
	wg.Wait()
}
func fetch(url string, ch chan string) {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if cancelled() {
		return
	}
	cancelChan := make(chan struct{})
	req.Cancel = cancelChan

	go func() {
		select {
		case <-done:
			close(cancelChan)
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case ch <- fmt.Sprint(err):
		case <-done:
		}
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	select {
	case ch <- fmt.Sprintf("%7d  %s", nbytes, url):
	case <-done:
	}
}
