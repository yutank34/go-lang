package ex16

import "bytes"

func join(sep string, strs ...string) string {
	var bf bytes.Buffer
	bf.WriteString(sep)
	for _, s := range strs {
		bf.WriteString(s)
	}
	return bf.String()
}
