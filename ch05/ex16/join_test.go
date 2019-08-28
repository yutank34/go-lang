package ex16

import "testing"

func TestJoin(t *testing.T) {
	var tests = []struct {
		sep  string
		strs []string
		want string
	}{
		{"a", []string{"b", "c", "d"}, "abcd"},
		{"", []string{"b", "c", "d"}, "bcd"},
		{"a", []string{}, "a"},
	}

	for _, test := range tests {
		output := join(test.sep, test.strs...)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}

}
