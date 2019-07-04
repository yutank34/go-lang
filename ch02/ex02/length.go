package main

import "fmt"

type Meters float64
type Feets float64

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feets) String() string  { return fmt.Sprintf("%gft", f) }

func MToF(m Meters) Feets { return Feets(m * 3.2808) }
func FToM(f Feets) Meters { return Meters(f / 3.2808) }
