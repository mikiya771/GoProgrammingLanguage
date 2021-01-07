package main

import (
	"fmt"
	"math"
	"sync"
)

type Counter struct {
	Count      int
	CountMutex *sync.Mutex
}

func main() {
	const workingCount = 10
	const workingNumber = 50
	var wg sync.WaitGroup
	for i := 0; i < workingCount; i++ {
		wg.Add(1)
		go func(num float64) {
			f := 1.0
			e := num
			for math.Abs(e) > 0.01 {
				e = f*f - num
				f -= e / (2 * f)
			}
			fmt.Println(f)
			wg.Done()
		}(workingNumber + float64(i))
	}
	wg.Wait()
}
