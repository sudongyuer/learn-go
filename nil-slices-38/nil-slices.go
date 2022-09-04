package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s)
	if s == nil {
		fmt.Println("nil!")
	}
}
