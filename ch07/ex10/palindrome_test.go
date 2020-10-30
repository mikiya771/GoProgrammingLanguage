package main

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for _, test := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", false},
		{"Evil I did dwell; lewd did I live.", false},
		{"Able was I ere I saw Elba", true},
		{"ete", true},
		{"Et se resservir, ivresse reste.", false},
		{"palindrome", false},
		{"desserts", false},
	} {
		if got := IsPalindrome(runes([]rune(test.input))); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
	for _, test := range []struct {
		input []string
		want  bool
	}{
		{[]string{"a", "a"}, true},
	} {
		if got := IsPalindrome(sort.StringSlice(test.input)); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
