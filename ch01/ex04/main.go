package main

import (
	"bufio"
	"fmt"
	"os"
)

type Files []string

func main() {
	filemap := make(map[string]Files)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No Files")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			fileLines(f, arg, filemap)
			f.Close()
		}
	}
	printedFiles := make(map[string]int)
	for _, fileNames := range filemap {
		for _, fn := range fileNames {
			if _, ok := printedFiles[fn]; !ok {
				printedFiles[fn]++
				fmt.Printf("%s\n", fn)
			}
		}
	}
}

func fileLines(f *os.File, name string, filemap map[string]Files) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		filemap[input.Text()] = append(filemap[input.Text()], name)
	}
}
