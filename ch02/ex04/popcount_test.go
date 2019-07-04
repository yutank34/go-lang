package popcount2

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

func TestPopCountV3(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{10, 2},
		{100, 3},
		{1000, 6},
	}
	for _, test := range tests {
		output := PopCountV3(test.input)
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

func BenchmarkPopCountV2(b *testing.B) {
	var tests = makeInputs()
	for _, v := range tests {
		PopCountV2(v)
	}
}

func BenchmarkPopCountV3(b *testing.B) {
	var tests = makeInputs()
	for _, v := range tests {
		PopCountV3(v)
	}
}
