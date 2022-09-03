package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v1 = Vertex{1, 2}
	var v2 = v1
	fmt.Println(Vertex{1, 2})
	fmt.Println(v1 == v2)
}
