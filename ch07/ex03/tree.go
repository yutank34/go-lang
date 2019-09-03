package ex03

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var trace func(*tree, int)
	bf := new(bytes.Buffer)

	trace = func(n *tree, depth int) {
		if n == nil {
			return
		}
		bf.WriteString(fmt.Sprintf("%*s%d\n", depth*2, "", n.value))
		trace(n.left, depth+1)
		trace(n.right, depth+1)
	}
	trace(t, 0)
	return bf.String()
}
