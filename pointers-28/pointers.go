package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	p := &i
	x := &p
	fmt.Println(p)
	fmt.Println(x)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
