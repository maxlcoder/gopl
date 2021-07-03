package main

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (f Foot) Sting() string {
	return fmt.Sprintf("%gft", f)
}

func FToM(f Foot) Meter {
	return Meter(f / 3.2808)
}

func MToF(m Meter) Foot {
	return Foot(m * 3.2808)
}
