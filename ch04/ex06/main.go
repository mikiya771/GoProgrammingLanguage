package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/labstack/gommon/log"
)

func main() {
	s := "hello   world"
	fmt.Println(s)
	s = string(convSpaces([]byte(s)))
	fmt.Println(s)
}
func convSpaces(b []byte) []byte {
	sb := make([]byte, 4)
	ss := utf8.EncodeRune(sb, ' ')
	space := sb[:ss]
	inSpace := false
	ip := 0
	var size int
	var r rune
	for next := 0; next < len(b); next += size {
		r, size = utf8.DecodeRune(b[next:])

		if r == utf8.RuneError {
			log.Error("RuneError")
			return b
		}

		if unicode.IsSpace(r) {
			if !inSpace {
				copy(b[ip:], space)
				ip += ss
				inSpace = true
			}
			continue
		}

		copy(b[ip:], b[next:next+size])
		ip += size
		inSpace = false
	}

	return b[:ip]
}
