package diffBitCount

import (
	"crypto/sha256"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func DiffPopCount(a, b string) int {
	c1 := sha256.Sum256([]byte(a))
	c2 := sha256.Sum256([]byte(b))

	var r int
	for i := 0; i < 32; i++ {
		c := c1[i] ^ c2[i]
		r += int(pc[c])
	}
	return r
}
