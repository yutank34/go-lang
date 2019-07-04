package popcount2

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
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>uint64(i*8))]
	}
	return int(result)
}

func PopCountV3(x uint64) int {
	v := 0
	for i := uint64(0); i < 64; i++ {
		v += int((x >> i) & 1)
	}
	return v
}
