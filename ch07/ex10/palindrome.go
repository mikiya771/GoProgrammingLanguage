package main

import (
	"sort"
	"unicode"
)

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}

type runes []rune

func (r runes) Len() int {
	return len(r)
}

func (r runes) Less(i, j int) bool {
	ri := r[i]
	rj := r[j]
	if unicode.IsLetter(ri) {
		ri = unicode.ToLower(ri)
	}
	if unicode.IsLetter(rj) {
		rj = unicode.ToLower(rj)
	}
	return ri < rj
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
