package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var total int64 = 0
	b := []byte("hello 世界")
	w := os.Stdout
	ww, c := CountingWriter(w)
	n, err := ww.Write(b)
	if err != nil {
		return
	}
	total += int64(len(b))
	fmt.Println(n, total, *c)
	n, err = ww.Write(b)
	if err != nil {
		return
	}
	total += int64(len(b))
	fmt.Println(n, total, *c)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var ww WriteWrapper
	ww.writer = w
	return &ww, &(ww.counter)
}

type WriteWrapper struct {
	writer  io.Writer
	counter int64
}

func (ww *WriteWrapper) Write(b []byte) (int, error) {
	n, err := ww.writer.Write(b)
	ww.counter += int64(n)
	return n, err
}
