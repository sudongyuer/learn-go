package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	list := List[int]{val: 1, next: &List[int]{val: 2}}
	fmt.Println(list)
}
