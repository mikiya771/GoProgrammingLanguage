package main

import (
	"unicode/utf8"
)

func reverseUTF8(s []byte) {
	revByteEachRune := func(s []byte) func() bool {
		pos := 0
		return func() bool {
			if pos >= len(s) {
				return false
			}
			r, size := utf8.DecodeRune(s[pos:])
			if r == utf8.RuneError {
				panic("the value is not utf8 valid byte")
			}
			reverseByte(s[pos : pos+size])
			pos += size
			return true
		}
	}
	isRemained := true
	f := revByteEachRune(s)
	for isRemained {
		isRemained = f()
	}
	reverseByte(s)
}
func reverseByte(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
