package ex12

import (
	"strconv"
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		input []string
		want  bool
	}{
		{[]string{"Christmas", "trims cash"}, true},
		{[]string{"Christmas", "trims a cash"}, false},
		{[]string{"Christmas", "trims cash!"}, true},
		{[]string{"Christmas", "trims cash?"}, true},
	}

	for _, test := range tests {
		output := anagram(test.input[0], test.input[1])
		if output != test.want {
			t.Error("unexpected value: expected=" + strconv.FormatBool(test.want) + " ,result=" + strconv.FormatBool(output))
		}
	}
}
