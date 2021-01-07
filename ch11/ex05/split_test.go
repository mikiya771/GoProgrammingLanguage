package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		str      string
		sep      string
		expected int
	}{
		{"a:b:c", ":", 3},
		{"a,b,c,d,e,f,g", ",", 7},
	}
	for _, test := range tests {
		words := strings.Split(test.str, test.sep)
		if len(words) != test.expected {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.str, test.sep, len(words), test.expected)
		}
	}

}
