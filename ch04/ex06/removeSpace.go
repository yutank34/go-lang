package ex06

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

}

func removeSpace(s []byte) ([]byte, bool) {
	i, bsize := 0, 0
	var result []byte
	var justBefore []byte

	for i < len(s) {
		if s[i]&(1<<7) == 0 {
			bsize = 1
		} else if s[i]&(1<<6) != 0 && s[i]&(1<<5) == 0 {
			bsize = 2
		} else if s[i]&(1<<6) != 0 && s[i]&(1<<5) != 0 && s[i]&(1<<4) == 0 {
			bsize = 3
		} else if s[i]&(1<<6) != 0 && s[i]&(1<<5) != 0 && s[i]&(1<<4) != 0 && s[i]&(1<<3) != 0 {
			bsize = 4
		} else {
			fmt.Print("error")
		}

		b := s[i : i+bsize]
		r, _ := utf8.DecodeRune(b)
		// 初期処理
		if len(justBefore) == 0 {
			if unicode.IsSpace(r) {
				result = append(result, ' ')
			} else {
				result = append(result, b...)
			}
			justBefore = b
			i = i + bsize
			continue
		}

		if unicode.IsSpace(r) {
			br, _ := utf8.DecodeRune(justBefore)
			if !unicode.IsSpace(br) {
				result = append(result, ' ')
			}
		} else {
			result = append(result, b...)
		}
		justBefore = b
		i = i + bsize
	}
	return result, true

}
