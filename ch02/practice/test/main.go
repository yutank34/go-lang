package main

import "fmt"

type Hoge struct {
	x, y int
}

func main() {
	// x := 1
	// p := &x
	// fmt.Println(p)
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(*p)

	// v := Hoge{1, 2}

	// pp := &v
	// fmt.Println(pp)
	// fmt.Println(*pp)
	x := f()
	fmt.Println(x)
	fmt.Println(&x)
	y := f()
	fmt.Println(y)
	fmt.Println(&y)
}

func f() int {
	v := 1
	fmt.Println("f: ")
	fmt.Println(&v)
	return v
}
