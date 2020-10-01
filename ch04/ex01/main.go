package main

import "crypto/sha256"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(1&i)
	}
}

func totalPopCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(pc[b])
	}
	return total
}

func calcDiff(a, b []byte) int {
	sha1 := sha256.Sum256(a)
	sha2 := sha256.Sum256(b)
	xorBytes := make([]byte, sha256.Size)
	for i := 0; i < sha256.Size; i++ {
		xorBytes[i] = sha1[i] ^ sha2[i]
	}
	return totalPopCount(xorBytes)
}
