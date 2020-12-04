package main

import (
	"fmt"
	"time"
)

func pipe(ch <-chan int, stages int, final chan<- int) {
	next := make(chan int)

	stages++
	if stages >= 4000000 {
		for v := range ch {
			final <- v
		}
	} else {
		go pipe(next, stages, final)

		for v := range ch {
			next <- v
		}
	}
}
func main() {
	start := time.Now()
	next := make(chan int)
	final := make(chan int)
	go pipe(next, 0, final)
	next <- 0
	fmt.Println(<-final)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
