package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fnames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fnames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, fnames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s: %s\n", n, line, strings.Join(fnames[line][0:], " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fnames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		f := f.Name()
		names := fnames[t]
		if contain(names, f) {
			continue
		}
		names = append(names, f)
		fnames[t] = names
	}
	// TODO: input.Err()のからのエラー処理
}

func contain(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}
