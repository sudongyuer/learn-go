package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Split(s, " ")
	res := make(map[string]int)
	for _, v := range arr {
		_, ok := res[v]
		if !ok {
			res[v] = 1
		} else {
			res[v]++
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
