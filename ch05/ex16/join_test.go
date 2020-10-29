package main

import (
	"testing"
)

var testCases = []struct {
	a        []string
	sep      string
	expected string
}{
	{nil, "", ""},
	{[]string{"abc"}, " ", "abc"},
	{[]string{"abc", "def", "ghi"}, ",", "abc,def,ghi"},
	{[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, " ",
		"0 1 2 3 4 5 6 7 8 9"},
}

func TestJoin(t *testing.T) {
	for _, tc := range testCases {
		result := join(tc.sep, tc.a...)
		if result != tc.expected {
			t.Errorf("[join(%s, %v)] return %s, want %s", tc.sep, tc.a, result, tc.expected)
		}
	}
}
