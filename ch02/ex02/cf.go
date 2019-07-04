package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

var types = flag.String("t", "temp", "type")

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		for _, arg := range flag.Args() {
			f, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			print(*types, f)
		}
	}
	cs := bufio.NewScanner(os.Stdin)
	for cs.Scan() {
		s := cs.Text()
		if s == "exit" {
			break
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("数字を入力してください")
		}
		print(*types, f)
	}
}

func print(t string, i float64) {
	switch t {
	case "temp":
		printTemp(i)
	case "length":
		printLength(i)
	case "weight":
		printTemp(i)
	default:
		fmt.Println("オプションが適切ではありません")
		os.Exit(1)
	}
}

func printTemp(a float64) {
	f := tempconv.Fahrenheit(a)
	c := tempconv.Celsius(a)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func printLength(a float64) {
	m := Meters(a)
	f := Feets(a)
	fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), f, FToM(f))
}

func printWeight(a float64) {
	k := Kilogram(a)
	p := Pound(a)
	fmt.Printf("%s = %s, %s = %s\n", k, KToP(k), p, PToK(p))
}
