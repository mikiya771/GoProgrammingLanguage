package main

import (
	"fmt"
	"testing"
)

const (
	ff = 1 << (8 * iota)
	ff2
	ff3
	ff4
	ff5
	ff6
	ff7
	ff8
)

var benchNum = []uint64{
	ff,
	ff2,
	ff3,
	ff4,
	ff5,
	ff6,
	ff7,
	ff8,
}
var fns = map[string]func(v uint64) int{
	"PopCount":            PopCount,
	"LoopedPopCount":      LoopedPopCount,
	"LongLoopedPopCount":  LongLoopedPopCount,
	"LogicCalcedPopCount": LogicCalcedPopCount,
}

func BenchmarkPopCount(b *testing.B) {
	for _, v := range benchNum {
		for k, vf := range fns {
			b.Run(fmt.Sprintf("function=%s, input-variable=%d", k, v), func(b *testing.B) {
				var s int
				for i := 0; i < b.N; i++ {
					s += vf(v)
				}
				fmt.Println(s)
			})
		}
	}
}
