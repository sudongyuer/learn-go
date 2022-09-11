package main

import "fmt"

type person struct {
	name    string
	friends map[int]string
}

func main() {
	p := person{name: "aaa"}
	p.friends = map[int]string{1: "bob"}
	fmt.Println(p)
	change(&p)
	fmt.Println(p)
}

func change(p *person) {
	p.name = "mmm"
	p.friends[1] = `lihua`
}
