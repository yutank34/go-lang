package ex07

import (
	"fmt"
	"os"
)

func reverse(s []byte) ([]byte, bool) {
	i, size := 0, 0
	result := make([]byte, len(s), cap(s))
	for i < len(s) {
		size = getSize(s[i])
		for j, k := len(s)-i-1, i+size-1; k >= i; j, k = j-1, k-1 {
			result[j] = s[k]
		}
		i = i + size
	}
	return result, true
}

func getSize(s byte) int {
	size := 0
	if s&(1<<7) == 0 {
		size = 1
	} else if s&(1<<6) != 0 && s&(1<<5) == 0 {
		size = 2
	} else if s&(1<<6) != 0 && s&(1<<5) != 0 && s&(1<<4) == 0 {
		size = 3
	} else if s&(1<<6) != 0 && s&(1<<5) != 0 && s&(1<<4) != 0 && s&(1<<3) != 0 {
		size = 4
	} else {
		fmt.Print("error")
		os.Exit(1)
	}
	return size
}
