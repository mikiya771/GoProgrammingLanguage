package main

import "testing"

var data = []struct {
	s1 string
	s2 string
	ex bool
}{
	{"HELLO", "hello", false},
	{"HELLO", "hellow", false},
	{"HELLO", "ellow", false},
	{"llohew", "hello", false},
	{"llohew", "helloww", false},
	{"llohe", "hello", true},
	{"hello", "hello", true},
	{"world hell", "hello wrld", true},
}

func TestIsAnagram(t *testing.T) {
	for _, d := range data {
		a := isAnagram(d.s1, d.s2)
		if a != d.ex {
			t.Errorf("%v is expected, but %v is returned by IsAnagram Function(%s, %s)", d.ex, a, d.s1, d.s2)
		}
	}
}
