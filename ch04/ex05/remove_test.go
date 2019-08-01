package ex05

import (
	"strconv"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{[]string{"aa", "aa", "bb"}, []string{"aa", "bb", ""}},
		{[]string{"aa", "aa", "aa", "bb"}, []string{"aa", "bb", "", ""}},
		{[]string{"aa", "bb"}, []string{"aa", "bb"}},
		{[]string{"aa", "bb", "bb"}, []string{"aa", "bb", ""}},
	}

	for _, test := range tests {
		removeDuplicate(test.input)
		if len(test.input) != len(test.want) {
			t.Error("Invalid length")
		}
		for i := 0; i < len(test.want); i++ {
			if test.input[i] != test.want[i] {
				t.Error("unexpected value: expected=" + test.want[i] + " ,result=" + test.input[i] + ", on" + strconv.Itoa(i))
			}
		}
	}
}
