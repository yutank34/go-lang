// 起動したコマンド名とそのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	Print(os.Stdout, os.Args)
}

func Print(w io.Writer, args []string) {
	fmt.Fprintf(w, "command: %s, args: %s\n", args[0], strings.Join(args[1:], " "))
}
