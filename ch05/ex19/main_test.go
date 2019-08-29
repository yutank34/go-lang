package ex19

import (
	"testing"
)

func TestF(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
	}

	for _, test := range tests {
		output := f(test.input)
		if output != test.input {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}
