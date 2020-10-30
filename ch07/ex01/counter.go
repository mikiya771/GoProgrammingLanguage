package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		_ = scanner.Text()
		*c += 1
	}
	return len(p), nil
}
func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	for scanner.Scan() {
		_ = scanner.Text()
		*c += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return len(p), nil

}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
