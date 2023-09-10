package main

import "fmt"

func fibonnaci() func() int {
	prev := -1
	curr := 1
	return func() int {
		num := prev + curr
		prev = curr
		curr = num
		return num
	}
}

func main() {
	f := fibonnaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
