package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("-", "dog", "cat", "mouse"))
	fmt.Println(join("-", "dog", "cat"))
	fmt.Println(join("-", "dog"))
}

func join(sep string, str ...string) string {
	return strings.Join(str, sep)
}
