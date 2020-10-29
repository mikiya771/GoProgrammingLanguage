package transpiler

import "testing"

func TestExpand(t *testing.T) {
	table := []struct {
		name    string
		raw     string
		exepect string
	}{
		{
			"No Expand",
			"animal fool bar fool",
			"animal fool bar fool",
		},
		{
			"Expand",
			"$animal fool bar fool",
			"cat fool bar fool",
		},
	}
	for _, o := range table {
		act := expand(o.raw, Transpiler)
		if act != o.exepect {
			t.Errorf("Test Failed %s, input: %s, output: %s, expect: %s", o.name, o.raw, act, o.exepect)
		}
	}
}
