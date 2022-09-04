package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5, 6}
	printSlice(s)
	arr1 := s[:3]
	arr1 = append(arr1, 10)
	printSlice(arr1)
	printSlice(s)
	// append works on nil slices.
	arr3 := append(s, 0)
	arr3[0] = 100
	printSlice(arr3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
