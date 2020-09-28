package main

import "fmt"

const (
	K  = 1000
	KB = K
	MB = KB * K
	GB = MB * K
	TB = GB * K
	PB = TB * K
	EB = PB * K
	ZB = EB * K
	YB = ZB * K
)

func main() {
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)
	fmt.Println("EB: ", EB)
}
