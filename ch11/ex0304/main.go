package main

import "unicode/utf8"

func IsPalindrome(l string) bool {
	ignoreTokens := map[rune]bool{
		',': true,
		'.': true,
		' ': true,
	}
	runes := []rune{}
	for len(l) > 0 {
		r, size := utf8.DecodeRuneInString(l)
		l = l[size:]
		if !ignoreTokens[r] {
			runes = append(runes, r)
		}
	}
	for i := 0; i < (len(runes)+1)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
