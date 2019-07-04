// コマンドライン引数のインデックスと値を1行ずつ表示する
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Print(os.Stdout, os.Args)
}

func Print(w io.Writer, args []string) {
	for i, arg := range args[1:] {
		fmt.Fprintf(w, "index: %d, arg: %s\n", i, arg)
	}
}
