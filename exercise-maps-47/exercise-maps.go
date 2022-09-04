package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Split(s, " ") {
		_, ok := m[w]
		if !ok {
			m[w] = 1
		} else {
			m[w] += 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
