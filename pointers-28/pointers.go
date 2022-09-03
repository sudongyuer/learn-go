package main

import (
	"fmt"
)

func main() {
	i, j, y := 42, 2701, 42
	arr1, arr2 := [3]int{1, 2, 3}, [3]int{1, 2, 3}
	p := &i
	x := &p
	c := &y
	parr1, parr2 := &arr1, &arr2
	fmt.Println(p)
	fmt.Println(x)
	fmt.Println(c == p)
	fmt.Println(parr1, parr2)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
