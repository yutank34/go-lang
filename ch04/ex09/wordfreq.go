package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

func wordfreq() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	for w, n := range counts {
		fmt.Printf("word:%v, count:%v\n", w, n)
	}

}
