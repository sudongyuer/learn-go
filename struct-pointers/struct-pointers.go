package main

import "fmt"

type Vertex struct {
	X int

	Y int
}

func main() {
	v := Vertex{1, 2}
	pointerv := &v
	num := &v
	pointerx := &num
	fmt.Println(num)
	fmt.Println(pointerx)
	fmt.Println(pointerv)
	pointerv.X = 4
	fmt.Println(v.X)
}
