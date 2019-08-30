package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadCast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadCast }
func IsCast(v Flags) bool   { return v&(FlagBroadCast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // 10001 true
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // 10000 false
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // 10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) // 10010 true
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB //1048576
	GiB //1073741824
	TiB
	PiB
	EiB
	ZiB
	YiB
)