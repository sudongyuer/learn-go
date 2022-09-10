package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	ch := make(chan int)
	go Walk(t1, ch)
	for i := range ch {
		fmt.Println(i)
		if i >= 10 {
			break
		}
	}
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
	t3 := tree.New(2)
	fmt.Println(Same(t1, t3))

}
