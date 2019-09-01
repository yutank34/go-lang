package ex02

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		input string
	}{
		{""},
		{"test is test"},
		{"テストです"},
	}

	for _, test := range tests {
		b := new(bytes.Buffer)
		w, c := CountingWriter(b)
		w.Write([]byte(test.input))

		if b.String() != test.input {
			t.Error("byteBufferに書き込めていない")
		}

		if *c != int64(len(test.input)) {
			t.Error("countが間違えている")
		}
	}
}
