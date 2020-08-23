package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}
func LoopedPopCount(x uint64) int {
	sum := byte(0)
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}
func LongLoopedPopCount(x uint64) int {
	sum := uint64(0)
	for i := 0; i < 64; i++ {
		sum += (x >> i) % 2
	}
	return int(sum)
}
func LogicCalcedPopCount(x uint64) int {
	s := x
	count := 0
	for s != 0 {
		count++
		s = s & (s - 1)
	}
	return count
}
