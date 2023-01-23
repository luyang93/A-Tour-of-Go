package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex8{3, 4}

	a = f // a MyFloat2 implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex8 implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex8 (not *Vertex8)
	// and does NOT implement Abser.
	//a = v
	//fmt.Println(a.Abs())
}
