package comma

import (
	"testing"
)

var data = [...]struct {
	input    string
	expected string
}{
	{"", ""},
	{"1", "1"},
	{"123", "123"},
	{"1234", "1,234"},
	{"123456", "123,456"},
	{"1234567", "1,234,567"},
	{"0123456789", "0,123,456,789"},
}

const N = 10000

func execute(t *testing.T, f func(string) string) {
	for _, d := range data {
		result := f(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}

func TestComma(t *testing.T) {
	execute(t, Comma)
}
func TestRepeatComma(t *testing.T) {
	execute(t, RepeatComma)
}
func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			Comma(d.input)
		}
	}
}

func BenchmarkCommaWithoutBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			RepeatComma(d.input)
		}
	}
}
