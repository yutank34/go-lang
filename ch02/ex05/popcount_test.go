package popcount3

import (
	"testing"
)

func TestPopCountV2(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{10, 2},
		{100, 3},
		{1000, 6},
	}
	for _, test := range tests {
		output := PopCountV2(test.input)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}

func TestPopCountV4(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{10, 2},
		{100, 3},
		{1000, 6},
	}
	for _, test := range tests {
		output := PopCountV4(test.input)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}

func makeInputs() []uint64 {
	var ints []uint64
	for i := 10000; i < 10000000; i++ {
		ints = append(ints, uint64(i))
	}
	return ints
}

var input uint64 = 111111
var output int

func BenchmarkPopCountV2(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountV2(input)
	}
	output = temp
}

func BenchmarkPopCountV3(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountV3(input)
	}
	output = temp
}

func BenchmarkPopCountV4(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountV4(input)
	}
	output = temp
}
