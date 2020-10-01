package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print(evaluateInput(os.Stdin))
}
func evaluateInput(input io.Reader) map[string]int {
	counts := make(map[string]int)
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	return counts
}
