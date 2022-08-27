package main

import "fmt"

func add(x int, y int) (z, g int) {
	z = x
	g = y
	return
}

func main() {
	fmt.Println(add(42, 13))
}
