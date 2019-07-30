package ex06

import (
	"strconv"
	"testing"
)

func TestRemoveSpace(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte("お腹空いた"), []byte("お腹空いた")},
		{[]byte("お腹 空いた"), []byte("お腹 空いた")},
		{[]byte("お腹  空いた"), []byte("お腹 空いた")},
		{[]byte("お腹\t\n空いた"), []byte("お腹 空いた")},
		{[]byte("お腹\v\f空いた"), []byte("お腹 空いた")},
		{[]byte("お腹\r 空いた"), []byte("お腹 空いた")},
		{[]byte("お腹\u00a0\u0085空いた"), []byte("お腹 空いた")},
		{[]byte("\t\nお腹空いた"), []byte(" お腹空いた")},
		{[]byte("\v\fお腹空いた"), []byte(" お腹空いた")},
		{[]byte("\r お腹空いた"), []byte(" お腹空いた")},
		{[]byte("\u00a0\u0085お腹空いた"), []byte(" お腹空いた")},
	}

	for _, test := range tests {
		output, _ := removeSpace(test.input)
		if len(output) != len(test.want) {
			t.Error("unexpected value: expectedLength=" + strconv.Itoa(len(test.want)) + " ,resultLength=" + strconv.Itoa(len(test.input)))
		}
		for i := 0; i < len(test.want); i++ {
			if output[i] != test.want[i] {
				t.Errorf("unexpected value: expected=%v ,result=%v, on%v", test.want[i], output[i], strconv.Itoa(i))
			}
		}
	}
}
