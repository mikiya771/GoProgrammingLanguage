package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("previous: ", s)
	rotate(&s, 2)
	fmt.Println("post: ", s)
}

func rotate(s *[]int, n int) {
	if n > len(*s) {
		fmt.Errorf("the input %d is over the length of %d", n, *s)
	}
	*s = append((*s)[n:], (*s)[:n]...)
}
