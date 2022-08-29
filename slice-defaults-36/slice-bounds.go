package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	var row = s
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	s[0] = 100
	fmt.Println(s)
	fmt.Println(row)
}
