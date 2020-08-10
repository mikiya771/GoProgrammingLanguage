package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(getCommandName())
	fmt.Println(s)
}

func getCommandName() string {
	return os.Args[0]
}
