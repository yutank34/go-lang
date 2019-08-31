package printints

import (
	"testing"
)

func TestIntsToString(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{[]int{1}, "[1]"},
		{[]int{1, 2}, "[1, 2]"},
		{[]int{1, 2, 3}, "[1, 2, 3]"},
	}

	for _, test := range tests {
		output := intsToString(test.input)

		if output != test.want {
			t.Error("unexpected value: expected=" + test.want + ",result=" + output)
		}
	}
}
