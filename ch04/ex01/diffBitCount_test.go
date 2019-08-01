package diffBitCount

import (
	"strconv"
	"testing"
)

func TestDiffBitCount(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{[]string{"x", "X"}, 125},
	}

	for _, test := range tests {
		output := DiffPopCount(test.input[0], test.input[1])
		if output != test.want {
			t.Error("unexpected value: expected=" + strconv.Itoa(test.want) + " ,result=" + strconv.Itoa(output))
		}
	}
}
