// Echo2 は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:]	{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
