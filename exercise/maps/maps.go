package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, value := range a {
		m[value] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
