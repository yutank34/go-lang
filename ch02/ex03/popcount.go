package popcount

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountV1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountV2(x uint64) int {
	var result byte
	for i := uint(0); i < 8; i++ {
		result += pc[byte(x>>i*8)]
	}
	return int(result)
}

func F() int {
	// fmt.Print(pc)
	var t [256]int
	for i := 0; i < 256; i++ {
		t[i] = int(i)
	}
	for i := 0; i < 256; i++ {
		fmt.Println((i & 1))
	}
	return int(1)
}
