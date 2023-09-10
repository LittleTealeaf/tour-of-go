package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)

	counts := make(map[string]int)

	for i := range fields {
		word := fields[i]
		counts[word] = counts[word] + 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
