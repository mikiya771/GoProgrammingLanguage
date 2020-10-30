package main

import "fmt"

type multikeySortTracks struct {
	tracks   []*Track
	lessFunc []lessFunc
}

func (ms *multikeySortTracks) Len() int { return len(ms.tracks) }
func (ms *multikeySortTracks) Less(i, j int) bool {
	if len(ms.lessFunc) == 0 {
		panic("No keys")
	}

	for ind := 0; ind < len(ms.lessFunc); ind++ {
		less := ms.lessFunc[ind]
		if less(ms.tracks[i], ms.tracks[j]) {
			return true
		}
		if less(ms.tracks[i], ms.tracks[j]) {
			return false
		}
	}
	return false
}
func (ms *multikeySortTracks) Swap(i, j int) { ms.tracks[i], ms.tracks[j] = ms.tracks[j], ms.tracks[i] }
func (ms *multikeySortTracks) AddSortKey(f lessFunc) {
	fmt.Println(f, ms.lessFunc)
	ms.lessFunc = append(ms.lessFunc, f)
	fmt.Println(f, ms.lessFunc)

}
