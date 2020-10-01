package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Print(evaluateInput(os.Stdin))
}
func evaluateInput(input io.Reader) (map[rune]int, map[string]int, int) {
	counts := make(map[rune]int)
	kinds := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(input)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v \n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			kinds["letter"]++
		}
		if unicode.IsDigit(r) {
			kinds["digit"]++
		}
		counts[r]++
		utflen[n]++
	}
	return counts, kinds, invalid
}
