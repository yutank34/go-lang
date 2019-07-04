package main

import (
	"math"
	"testing"
)

const ev float64 = 0.00000001

func TestKToP(t *testing.T) {
	var tests = []struct {
		input Kilogram
		want  Pound
	}{
		{0, 0},
		{100, 220.46},
	}

	for _, test := range tests {
		output := KToP(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestPToK(t *testing.T) {
	var tests = []struct {
		input Pound
		want  Kilogram
	}{
		{0, 0},
		{220.46, 100},
	}

	for _, test := range tests {
		output := PToK(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestMToF(t *testing.T) {
	var tests = []struct {
		input Meters
		want  Feets
	}{
		{0, 0},
		{100, 328.08},
	}

	for _, test := range tests {
		output := MToF(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}

func TestFToM(t *testing.T) {
	var tests = []struct {
		input Feets
		want  Meters
	}{
		{0, 0},
		{328.08, 100},
	}

	for _, test := range tests {
		output := FToM(test.input)
		if math.Abs(float64(output-test.want)) > ev {
			t.Errorf("unexpected value: expected: %f result: %f", test.want, output)
		}
	}
}
