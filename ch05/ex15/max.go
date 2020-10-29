package variadic

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("len of input vals is zero")
	}
	ret := vals[0]
	for _, v := range vals {
		if v > ret {
			ret = v
		}
	}
	return ret, nil
}
func maxWithALO(fv int, vals ...int) int {
	ret := fv
	for _, v := range vals {
		if v > ret {
			ret = v
		}
	}
	return ret
}
