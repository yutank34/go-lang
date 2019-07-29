package ex11

import "testing"

func TestCommaHandler(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"100", "100"},
		{"1000", "1,000"},
		{"10000", "10,000"},
		{"100000", "100,000"},
		{"1000000", "1,000,000"},
		{"0.1", "0.1"},
		{"10.01", "10.01"},
		{"100.001", "100.001"},
		{"1000.0001", "1,000.0001"},
	}

	for _, test := range tests {
		output := commaHandler(test.input)
		if output != test.want {
			t.Error("unexpected value: expected=" + test.want + " ,result=" + output)
		}
	}
}

func TestFloatComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"0.1", "0.1"},
		{"10.01", "10.01"},
		{"100.001", "100.001"},
		{"1000.0001", "1,000.0001"},
	}

	for _, test := range tests {
		output := floatComma(test.input)
		if output != test.want {
			t.Error("unexpected value: expected=" + test.want + " ,result=" + output)
		}
	}
}
