package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
