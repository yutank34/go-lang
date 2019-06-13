package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"command", ""}, "command: command, args: "},
		{[]string{"command", "h"}, "command: command, args: h"},
		{[]string{"command", "hoge"}, "command: command, args: hoge"},
		{[]string{"command", "hoge", "fuga"}, "command: command, args: hoge fuga"},
	}

	for _, test := range tests {
		buf := new(bytes.Buffer)
		Print(buf, test.input)
		output := strings.Trim(string(buf.Bytes()), "\n\t")

		if output != test.want {
			t.Error("unexpected value:\nexpected: " + test.want + "\nresult: " + output)
		}
	}
}
