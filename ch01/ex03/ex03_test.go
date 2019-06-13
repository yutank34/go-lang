package ex03

import (
	"testing"
)

func MakeStrings() []string {
	var words []string
	for i := 0; i < 100; i++ {
		words = append(words, "hoge")
	}
	return words
}

func BenchmarkPrint1(b *testing.B) {
	var test = MakeStrings()
	for i := 0; i < b.N; i++ {
		JoinWithStrings(test)
	}
}

func BenchmarkPrint2(b *testing.B) {
	var test = MakeStrings()
	for i := 0; i < b.N; i++ {
		JoinWithFor(test)
	}
}
