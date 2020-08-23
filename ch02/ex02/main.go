package main

import (
	"./length"
	"./weight"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := os.Args[1:]
	fl := len(inputs)
	for i := 0; i < 2-fl; i++ {
		fmt.Println("Please Input Number and Type(weight/length are valid types)")
		var s string
		fmt.Scan(&s)
		inputs = append(inputs, s)
	}
	num, err := strconv.ParseFloat(inputs[0], 64)
	if err != nil {
		os.Exit(1)
	}
	if inputs[1] == "weight" {
		fmt.Printf("%g g = %g lb, %g lb = %g g", num, weight.GtoP(weight.Gram(num)), num, weight.PtoG(weight.Pond(num)))
	} else if inputs[1] == "length" {
		fmt.Printf("%g m = %g in, %g in = %g m", num, length.MtoI(length.Meter(num)), num, length.ItoM(length.Inch(num)))
	} else {
		fmt.Printf("Sorry, type: %s is not valid now\n", inputs[1])
	}

}
