package main

import "io"

type nbyteReader struct {
	r     io.Reader
	limit int64
	next  int
}

func (nr *nbyteReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if int64(nr.next) >= nr.limit {
		return 0, io.EOF
	}
	nBytes := int((nr.limit - int64(nr.next)))
	if nBytes > len(p) {
		nBytes = len(p)
	}
	n, err := nr.r.Read(p[:nBytes])
	if err != nil {
		return n, err
	}
	nr.next += nBytes
	return n, nil
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &nbyteReader{r, n, 0}
}
