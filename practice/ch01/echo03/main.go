// Echo3 は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Paint1(w io.Writer, args []string) {
	fmt.Fprintf(w, strings.Join(args[1:], " "))
}

func Paint2(w io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintf(w, s)
}
