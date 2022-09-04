package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr1 := test(arr)
	fmt.Println(arr1)
	fmt.Println(arr)
}

func test(arr [5]int) [5]int {
	arr[0] = 100
	return arr
}
