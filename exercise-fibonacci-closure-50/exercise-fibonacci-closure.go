package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1 := 0
	fib2 := 1
	return func() int {
		fib2 = fib1 + fib2 //2
		fib1 = fib2 - fib1 //1
		return fib2 - fib1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
