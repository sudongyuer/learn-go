package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Channels Tutorial")
	values := make(chan int)
	time.Sleep(1000 * time.Millisecond)
}
