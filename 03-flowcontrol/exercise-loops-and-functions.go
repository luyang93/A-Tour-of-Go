package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value= %v\n", i, z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := x / 2
	for t := 0.0; math.Abs(z-t) > 1e-6; {
		t = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Value= %v\n", z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt")
	fmt.Println(Sqrt(100))
	fmt.Println("Sqrt2")
	fmt.Println(Sqrt2(100))
}
