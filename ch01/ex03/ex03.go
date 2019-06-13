// Echo3 は、そのコマンドライン引数を表示します。
package ex03

import (
	"strings"
)

func JoinWithStrings(args []string) string {
	return strings.Join(args[1:], " ")
}

func JoinWithFor(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
