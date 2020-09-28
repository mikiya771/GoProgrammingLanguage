package comma

import (
	"fmt"
	"strings"
)

func Comma(s string) string {
	n := len(s)
	if n == 0{
		return s
	}
	if s[0:1] == "+" || s[0:1] == "-"{
		return s[0:1] + Comma(s[1:])
	}
	sp := strings.Split(s, ".")
	if len(sp) > 2{
		fmt.Errorf("invalid input string")
		return s
	}
	if len(sp) == 2 {
		return strings.Join([]string{Comma(sp[0]),sp[1]}, ".")
	}
	if n <= 3{
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}
