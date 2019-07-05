package tempconv

import (
	"math"
	"testing"
)

const ev float64 = 0.00000001

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

		if math.Abs(float64(output-test.want)) > ev {
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

		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{-273, 0},
		{0, 273},
		{100, 373},
	}

	for _, test := range tests {
		output := CToK(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{0, -273},
		{273, 0},
		{373, 100},
	}

	for _, test := range tests {
		output := KToC(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}
