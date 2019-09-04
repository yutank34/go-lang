package ex04

import()

type StringReader struct {
	src string
	offset int
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	for {
		if r.offset + n >= len(r.src) {
			r.offset += n
			return n, io.EOF
		}

		if n >= len(p) {
			r.offset += n
			return n, nil
		}

		p[n] = r.src[r.offset+n]
		n++
	}
}

func NewReader(s string) io.Reader {
	return &StringReader(src: s)
}