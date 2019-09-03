package ex03

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		root *tree
		want string
	}{
		{
			&tree{value: 0},
			`0
`,
		},
		{
			&tree{
				0,
				&tree{value: 1},
				&tree{value: 2},
			},
			`0
  1
  2
`,
		},
	}
	for _, test := range tests {
		output := test.root.String()
		//		fmt.Println(output)
		//		fmt.Println(test.want)
		if output != test.want {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, output)
		}
	}
}
