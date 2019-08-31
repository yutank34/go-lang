package printints

import (
	"bytes"
	"fmt"
)

// intsToStringはfmt.Sprint(values)に似てるがカンマを追加してる
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
