package main

import "fmt"

func main() {
	a := make([]float64, 5)
	printSlice("a", a)

	b := make([]float64, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]

	printSlice("d", d)
	printSlice("b", b)
}

func printSlice(s string, x []float64) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
