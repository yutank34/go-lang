package main

import (
	"fmt"
)

const (
	KB = 1000      // 1e3
	MB = KB * 1000 // 1e6
	GB = MB * 1000 // 1e9
	TB = GB * 1000 // 1e12
	PB = TB * 1000 // 1e15
	EB = PB * 1000 // 1e18
	ZB = EB * 1000 // 1e21
	YB = ZB * 1000 // 1e24
)

func main() {
	valuess := []int64{KB, MB, GB, TB, PB, EB}
	for _, v := range valuess {
		fmt.Println(v)
	}
}
