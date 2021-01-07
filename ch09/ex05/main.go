package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Count      int
	CountMutex *sync.Mutex
}

func main() {
	const workingTime = 10
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	closer := make(chan struct{})
	counter := Counter{
		Count:      0,
		CountMutex: &sync.Mutex{},
	}
	go messagePublisher(ch1, ch2, "cat", closer, &counter)
	go messagePublisher(ch2, ch1, "dog", closer, &counter)
	ch1 <- struct{}{}
	<-time.Tick(time.Second * workingTime)
	close(closer)
	fmt.Printf("%d req/s", counter.Count/workingTime)

}
func messagePublisher(chin <-chan struct{}, chout chan<- struct{}, name string, closer <-chan struct{}, c *Counter) {
	for {
		select {
		case _ = <-chin:
			select {
			case chout <- struct{}{}:
				c.CountMutex.Lock()
				c.Count++
				c.CountMutex.Unlock()
			case <-closer:
				return
			}
		case <-closer:
			return
		}
	}
}
