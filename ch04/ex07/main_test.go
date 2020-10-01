package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	td := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"おはよう", "うよはお"},
	}
	for _, d := range td {
		b := []byte(d.input)
		if reverseUTF8(b); string(b) != d.expected {
			t.Errorf("uncorrect reverse....\n Input:%s\n, Expect:%s\n, Actual:%s\n", d.input, d.expected, string(b))
		}
	}
}
