package comma

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"100", "100"},
		{"100000", "100,000"},
		{"1000000", "1,000,000"},
	}

	for _, test := range tests {
		output := comma(test.input)
		if output != test.want {
			t.Error("unexpected value: expected=" + test.want + ", result=" + output)
		}
	}
}
