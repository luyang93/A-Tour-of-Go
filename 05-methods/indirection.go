package main

import "fmt"

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex5{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex5{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
