package main

import "testing"

var testdata = []struct {
	st1      []byte
	st2      []byte
	expected int
}{
	{[]byte("x"), []byte("x"), 0},
	{[]byte("X"), []byte("X"), 0},
	{[]byte("x"), []byte("X"), 125},
}

func TestCalcDiff(t *testing.T) {
	for _, d := range testdata {
		result := calcDiff(d.st1, d.st2)
		if result != d.expected {
			t.Errorf("Result of sha256 diff is wrong. %d is expected, but actual is %d", d.expected, result)
		}
	}
}
