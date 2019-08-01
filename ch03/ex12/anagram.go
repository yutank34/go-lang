package ex12

import (
	"sort"
	"strings"
)

var symbols = []string{" ", ",", ".", "!", "?"}

func anagram(s1, s2 string) bool {
	// 大文字、小文字は区別しない
	// 記号、空白は無視する
	s1 = removeSymbol(s1)
	s2 = removeSymbol(s2)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	return compareByts(b1, b2)
}

func compareByts(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func removeSymbol(s string) string {
	for _, symbol := range symbols {
		s = strings.ReplaceAll(s, symbol, "")
	}
	return s
}
