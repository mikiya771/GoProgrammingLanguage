package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		values []int
	}{
		{[]int{1}},
		{[]int{1, 144, 9, 42}},
		{[]int{1, 32, 32 << 1, 32 << 2, 32 << 3, 32 << 8}},
	} {
		var x IntSet
		m := map[int]bool{}
		for _, v := range tc.values {
			x.Add(v)
			m[v] = true
		}
		if x.Len() != len(m) {
			t.Errorf("x.Len() is %d, but want %v", x.Len(), len(m))
		}
	}

	var x IntSet
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}
}

func TestElems(t *testing.T) {
	for _, tc := range []struct {
		t []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 10, 100, 1000, 10000}},
		{[]int{1, 1 << 4, 1 << 6, 1 << 8, 1 << 10, 1 << 12, 1 << 14, 1 << 16}},
	} {
		var x IntSet
		m := map[int]bool{}
		x.AddAll(tc.t...)
		for _, v := range tc.t {
			m[v] = true
		}

		enums := x.Elems()
		if len(enums) != len(tc.t) {
			t.Errorf("len(enums) is %d, but want %d", len(enums), len(m))
			t.Errorf("enums is %v, t is %v", enums, tc.t)
		}

		for _, value := range enums {
			if _, ok := m[value]; !ok {
				t.Errorf("%d is not set to expected map, but want to be set", value)
			}
		}
		for value, _ := range m {
			if !x.Has(value) {
				t.Errorf("x.Has(%d) is false, but want true", value)
			}
		}
	}
}
