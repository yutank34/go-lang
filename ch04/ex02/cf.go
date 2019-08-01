package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var isSha256 = flag.Bool("2", false, "sha256")
var isSha384 = flag.Bool("3", false, "sha384")
var isSha512 = flag.Bool("5", false, "sha512")

func main() {
	flag.Parse()
	cs := bufio.NewScanner(os.Stdin)
	for cs.Scan() {
		s := cs.Text()
		if s == "exit" {
			break
		}

		if *isSha256 || (!*isSha256 && !*isSha384 && !*isSha512) {
			fmt.Printf("SHA-256 : %x\n", sha256.Sum256([]byte(s)))
		}

		if *isSha384 {
			fmt.Printf("SHA-384 : %x\n", sha512.Sum384([]byte(s)))
		}

		if *isSha512 {
			fmt.Printf("SHA-256 : %x\n", sha512.Sum512([]byte(s)))
		}
	}
	os.Exit(0)
}
