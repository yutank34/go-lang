package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"command", ""}, "index: 0, arg: "},
		{[]string{"command", "h"}, "index: 0, arg: h"},
		{[]string{"command", "hoge"}, "index: 0, arg: hoge"},
		{[]string{"command", "hoge", "fuga"}, "index: 0, arg: hoge\nindex: 1, arg: fuga"},
	}

	for _, test := range tests {
		buf := new(bytes.Buffer)
		Print(buf, test.input)
		output := strings.Trim(string(buf.Bytes()), "\n\t")

		fmt.Println(output)
		if output != test.want {
			t.Error("unexpected value:\nexpected:" + test.want + "\nresult:" + output)
		}
	}
}
