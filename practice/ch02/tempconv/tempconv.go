package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(BoilingC - FreezingC)
	boilinfF := CToF(BoilingC)
	fmt.Println(boilinfF - CToF(FreezingC))
	// fmt.Println(boilinfF - FreezingC)
	fmt.Println("========")
	fmt.Println(c == 0)
	fmt.Println(c > 0)
	fmt.Println(c)
	fmt.Println(f)
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius(f - 32*5/9) }

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
