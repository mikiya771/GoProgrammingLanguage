package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", 64+i+j)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		cnt += bits.OnesCount64(word)
	}
	return cnt
}
func (s *IntSet) Elems() []int {
	l := s.Len()
	if l == 0 {
		return []int{}
	}
	ret := make([]int, 0, l)
	for i, word := range s.words {
		for bit := uint(0); bit < 64; bit++ {
			if word&(1<<bit) != 0 {
				ret = append(ret, 64*i+int(bit))
			}
		}
	}
	return ret
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	t := []uint64{}
	_ = copy(t, s.words)
	n := &IntSet{
		words: t,
	}
	return n
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}
