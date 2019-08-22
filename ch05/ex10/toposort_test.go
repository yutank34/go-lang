package toposort

import "testing"

func TestTopoSort(t *testing.T) {
	tests := []struct {
		m    map[string][]string
		want []string
	}{
		{
			map[string][]string{
				"2":  []string{"11"},
				"9":  []string{"11", "8"},
				"10": []string{"11", "3"},
				"11": []string{"7", "5"},
				"8":  []string{"7", "3"},
			},
			[]string{"7", "5", "11", "3", "10", "2", "8", "9"},
		},
	}

	for _, test := range tests {
		outputs := topoSort(test.m)
		faild := false
		for i := 0; i < len(outputs); i++ {
			if outputs[i] != test.want[i] {
				faild = true
			}
		}
		if faild {
			t.Errorf("unexpected value:\nexpected: %v, result: %v", test.want, outputs)
		}
	}
}
