package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Fahrenheit
	}{
		{-273, -459.4},
		{0, 32},
		{100, 212},
	}

	for _, test := range tests {
		output := CToF(test.input)

		if output != test.want {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Celsius
	}{
		{-459.4, -273},
		{32, 0},
		{212, 100},
	}

	for _, test := range tests {
		output := FToC(test.input)

		if output != test.want {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{-273, 0.16},
		{0, 273.16},
		{100, 373.16},
	}

	for _, test := range tests {
		output := CToK(test.input)
		if output == test.want {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{0.16, -273},
		{273.16, 0},
		{373.16, 100},
	}

	for _, test := range tests {
		output := KToC(test.input)
		if output == test.want {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}
