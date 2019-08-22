package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(exptand("aaa aaa $hoge aaa", addS))
}

func exptand(s string, f func(string) string) string {
	var resultBf, inWordBf bytes.Buffer
	isTarget := false
	for _, r := range s {
		if r == '$' {
			isTarget = true
		} else if unicode.IsSpace(r) {
			isTarget = false
		}

		if isTarget {
			inWordBf.WriteRune(r)
		} else {
			if inWordBf.Len() != 0 {
				resultBf.WriteString(f(inWordBf.String()[1:]))
				inWordBf.Reset()
			}
			resultBf.WriteRune(r)
		}
	}
	return resultBf.String()
}

func addS(s string) string {
	return s + "S"
}
