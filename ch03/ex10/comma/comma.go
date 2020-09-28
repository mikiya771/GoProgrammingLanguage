package comma

import (
	"bytes"
	"fmt"
)

func Comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i, v := range s {
		if i != 0 && (n-i)%3 == 0 {
			fmt.Fprintf(&buf, "%s", ",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
func RepeatComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return RepeatComma(s[:n-3]) + "," + s[n-3:]
}
