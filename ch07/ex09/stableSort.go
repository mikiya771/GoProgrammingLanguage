package main

import "fmt"

type multikeyStableSortTracks struct {
	tracks   []*Track
	lessFunc []lessFunc
	keyPos   int
}
type lessFunc func(i, j interface{}) bool

func (ms multikeyStableSortTracks) Len() int { return len(ms.tracks) }
func (ms multikeyStableSortTracks) Less(i, j int) bool {
	if ms.keyPos < 0 {
		panic(fmt.Errorf("Out of Index: %d", ms.keyPos))
	}
	return ms.lessFunc[ms.keyPos](ms.tracks[i], ms.tracks[j])
}

func (ms multikeyStableSortTracks) HasNext() {
	ms.keyPos--
}
func (ms multikeyStableSortTracks) Reset() {
	ms.keyPos = len(ms.lessFunc)
}
func (ms multikeyStableSortTracks) AddLessFunc(lf lessFunc) {
	ms.lessFunc = append(ms.lessFunc, lf)
	ms.keyPos++
}
func (ms multikeyStableSortTracks) Swap(i, j int) {
	ms.tracks[i], ms.tracks[j] = ms.tracks[j], ms.tracks[i]
}
