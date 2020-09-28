package main

func isAnagram(a, b string) bool {
	aMap := map[rune]int{}
	bMap := map[rune]int{}
	for _, c := range a {
		aMap[c]++
	}
	for _, c := range b {
		bMap[c]++
	}
	for k, v := range aMap {
		if bMap[k] != v {
			return false
		}
	}
	for k, v := range bMap {
		if aMap[k] != v {
			return false
		}
	}
	return true
}
