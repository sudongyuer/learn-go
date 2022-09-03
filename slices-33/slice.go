package main

import "fmt"

func fn() {
	primes := [6]int{2, 3, 4, 5, 6, 7}
	var s []int = primes[0:3]
	fmt.Println(s, len(s), cap(s))

	s = s[:6]
	fmt.Println(s, len(s), cap(s))
}

func main() {
	// var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var arr2 = arr1
	// arr2[0] = 100
	// fmt.Println(arr, len(arr), cap(arr))
	// fmt.Println(arr1, len(arr1), cap(arr1))
	// fmt.Println(arr2, len(arr2), cap(arr2))
	// fmt.Println(arr1 == arr2)
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, len(slice), cap(slice))
	var slice1 = slice[0:3]
	slice1[0] = 100
	fmt.Println(slice1, len(slice1), cap(slice1))
	var slice2 = slice[3:6]
	slice2[0] = 200
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice, len(slice), cap(slice))
}
