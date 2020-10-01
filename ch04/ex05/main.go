package main

import "fmt"

func main() {
	s := []int{3, 5, 7, 7, 1, 1, 1}
	fmt.Println(s)
	s = removeAdjacentDuplicateElement(s)
	fmt.Println(s)
}

func removeAdjacentDuplicateElement(s []int) []int {
	i := 0
	for j := 1; j < len(s); j++ {
		if s[i] != s[j] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
