package main

import (
	"fmt"
	"testing"
)

var s int

func BenchmarkPopCount(b *testing.B) {
	n := uint64(1000000)
	b.ResetTimer()
	for i := uint64(0); i < n; i++ {
		s += PopCount(n * 13)
	}
	fmt.Println(s)
}

func BenchmarkLoopedPopCount(b *testing.B) {
	n := uint64(1000000)
	b.ResetTimer()
	for i := uint64(0); i < n; i++ {
		s += LoopedPopCount(n * 13)
	}
	fmt.Println(s)
}
func BenchmarkLongLoopedPopCount(b *testing.B) {
	n := uint64(1000000)
	b.ResetTimer()
	for i := uint64(0); i < n; i++ {
		s += LongLoopedPopCount(n * 13)
	}
	fmt.Println(s)
}
func BenchmarkCalkedPopCount(b *testing.B) {
	n := uint64(1000000)
	b.ResetTimer()
	for i := uint64(0); i < n; i++ {
		s += LogicCalcedPopCount(n * 13)
	}
	fmt.Println(s)
}
