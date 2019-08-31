package ex01

import "testing"

func TestLineCounter(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"test", 1},
		{"ああ\n", 1},
		{"ああ\nああ", 2},
	}

	for _, test := range tests {
		c := LineCounter(0)
		c.Write([]byte(test.input))
		if int(c) != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, c)
		}
	}
}

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"test", 4},
		{"test is test", 10},
		{"test は test ", 9},
	}

	for _, test := range tests {
		c := WordCounter(0)
		c.Write([]byte(test.input))
		if int(c) != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, c)
		}
	}
}
