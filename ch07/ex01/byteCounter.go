package ex01

import (
	"unicode"
	"unicode/utf8"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	*c++
	var isNewLine bool
	for _, b := range p {
		isNewLine = false
		if b == '\n' {
			*c++
			isNewLine = true
		}
	}
	if isNewLine {
		*c--
	}
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	for i, width := 0, 0; i < len(p); i += width {
		var r rune
		r, width = utf8.DecodeRune(p[i:])
		if !unicode.IsSpace(r) {
			*c++
		}
	}
	return len(p), nil
}
