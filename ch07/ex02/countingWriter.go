package ex02

import "io"

type byteCounter struct {
	writer    io.Writer
	byteCount int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.byteCount += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := byteCounter{w, 0}
	return &c, &(c.byteCount)
}
