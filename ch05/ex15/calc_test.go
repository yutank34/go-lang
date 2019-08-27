package ex15

import "testing"

func TestMax(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
		ok    bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{10, 2, 3, 4, 5}, 10, true},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		output, ok := max(test.input...)
		if test.ok != ok || output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
		ok    bool
	}{
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{10, 2, 3, 4, 5}, 2, true},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		output, ok := min(test.input...)
		if test.ok != ok || output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}

func TestMax2(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{10, 2, 3, 4, 5}, 10},
	}

	for _, test := range tests {
		output := max2(test.input[0], test.input[1:]...)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{10, 2, 3, 4, 5}, 2},
	}

	for _, test := range tests {
		output := min2(test.input[0], test.input[1:]...)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}
