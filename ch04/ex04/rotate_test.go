package ex04

import (
	"strconv"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		input []int
		step  int
		want  []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 0, []int{0, 1, 2, 3, 4, 5}},
		{[]int{0, 1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5, 0}},
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
		{[]int{0, 1, 2, 3, 4, 5}, 5, []int{5, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4, 5}, 7, []int{1, 2, 3, 4, 5, 0}},
	}

	for _, test := range tests {
		rotate(test.input, test.step)

		if len(test.input) != len(test.want) {
			t.Error("unexpected value: expectedLength=" + strconv.Itoa(len(test.want)) + " ,resultLength=" + strconv.Itoa(len(test.input)) + "on Step:" + strconv.Itoa(test.step))
		}
		for i := 0; i < len(test.want); i++ {
			if test.input[i] != test.want[i] {
				t.Error("unexpected value: expected=" + strconv.Itoa(test.want[i]) + " ,result=" + strconv.Itoa(test.input[i]) + ", on" + strconv.Itoa(i) + "on Step:" + strconv.Itoa(test.step))
			}
		}
	}
}
