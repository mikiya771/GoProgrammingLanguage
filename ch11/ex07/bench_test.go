package intset

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

var seed int64 = time.Now().UTC().UnixNano()

func BenchmarkAdd_IntSet(b *testing.B) {
	s := 0
	for i := 0; i < b.N; i++ {
		var set IntSet
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			set.Add(rng.Intn(math.MaxInt16))
		}
		s += set.Len()
	}
	fmt.Println(s)
}

func BenchmarkAdd_Map(b *testing.B) {
	s := 0
	for i := 0; i < b.N; i++ {
		m := make(map[int]bool)
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			m[rng.Intn(math.MaxInt16)] = true
		}
		s += len(m)
	}
	fmt.Println(s)
}

func BenchmarkUnionWith_IntSet(b *testing.B) {
	s := 0
	for i := 0; i < b.N; i++ {
		var x IntSet
		var y IntSet

		rng := rand.New(rand.NewSource(seed))

		for i := 0; i < 500; i++ {
			x.Add(rng.Intn(math.MaxInt16))
			y.Add(rng.Intn(math.MaxInt16))
		}
		x.UnionWith(&y)
		s += x.Len()
	}
	fmt.Println(s)
}

func BenchmarkUnionWith_IntMap(b *testing.B) {
	s := 0
	for i := 0; i < b.N; i++ {
		x := make(map[int]bool)
		y := make(map[int]bool)

		rng := rand.New(rand.NewSource(seed))

		for i := 0; i < 500; i++ {
			x[rng.Intn(math.MaxInt16)] = true
			y[rng.Intn(math.MaxInt16)] = true
		}
		for k, v := range y {
			x[k] = v
		}
		s += len(x)
	}
	fmt.Println(s)
}
