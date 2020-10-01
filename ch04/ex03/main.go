package main

import "fmt"

const (
	size = 3
)

func main() {
	a := [size]int{1, 2, 3}
	reverseArr(&a)
	fmt.Println(a)
}
func reverseArr(s *[size]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
