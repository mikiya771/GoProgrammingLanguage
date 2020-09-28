package comma

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
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
