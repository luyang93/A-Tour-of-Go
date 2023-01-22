package main

import "fmt"

func main() {
	var p1 *int
	fmt.Printf("Type %T, Value %v", p1, p1)

	i, j := 24, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
