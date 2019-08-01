package ex07

import (
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte("お腹すいた"), []byte("たいす腹お")},
	}

	for _, test := range tests {
		output, ok := reverse(test.input)

		if !ok {
			t.Errorf("Invalid Logic")
		}

		if len(output) != len(test.want) {
			t.Error("unexpected value: expectedLength=" + strconv.Itoa(len(test.want)) + " ,resultLength=" + strconv.Itoa(len(output)))
		}
		for i := 0; i < len(test.want); i++ {
			if output[i] != test.want[i] {
				t.Errorf("unexpected value: expected=%v ,result=%v, on%v", test.want[i], output[i], strconv.Itoa(i))
			}
		}
	}
}
