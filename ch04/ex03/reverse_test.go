package ex03

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input [5]int
		want  [5]int
	}{
		{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		reverse(&test.input)

		if test.input != test.want {
			t.Errorf("unexpected value: expected=%v, result=%v", test.want, test.input)
		}

		// if len(test.input) != len(test.want) {
		// 	t.Error("unexpected value: expectedLength=" + strconv.Itoa(len(test.want)) + " ,resultLength=" + strconv.Itoa(len(test.input)))
		// }
		// for i := 0; i < len(test.want); i++ {
		// 	if test.input[i] != test.want[i] {
		// 		t.Error("unexpected value: expected=" + strconv.Itoa(test.want[i]) + " ,result=" + strconv.Itoa(test.input[i]) + ", on" + strconv.Itoa(i))
		// 	}
		// }
	}
}
