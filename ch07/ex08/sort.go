package main

type multikeySortTracks struct {
	tracks   []*Track
	lessFunc []lessFunc
}
type lessFunc func(i, j interface{}) bool

func (ms multikeySortTracks) Len() int { return len(ms.tracks) }
func (ms multikeySortTracks) Less(i, j int) bool {
	if len(ms.lessFunc) == 0 {
		panic("No keys")
	}

	for ind := 0; ind < len(lessFunc); i++ {
		less := ms.lessFunc[ind]
		if less(i, j) {
			return true
		}
		if less(j, i) {
			return false
		}
	}
}
func (ms multikeySortTracks) Swap(i, j int) { ms.tracks[i], ms.tracks[j] = ms.tracks[j], ms.tracks[i] }
