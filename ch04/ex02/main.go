package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var shatype = flag.Int("sha", 256, "256(default), 384, 512")

func main() {
	flag.Parse()
	switch *shatype {
	case 256, 384, 512:
	default:
		fmt.Println("invalid sha", *shatype)
		flag.PrintDefaults()
		os.Exit(1)
	}
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("invalid arg length")
		os.Exit(1)
	}

	data := []byte(args[0])
	switch *shatype {
	case 256:
		sum := sha256.Sum256(data)
		fmt.Printf("%x\n", sum)
	case 384:
		sum := sha512.Sum384(data)
		fmt.Printf("%x\n", sum)
	case 512:
		sum := sha512.Sum512(data)
		fmt.Printf("%x\n", sum)
	}
}
