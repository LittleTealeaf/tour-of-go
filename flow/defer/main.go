package main

import "fmt"


func defer_stack() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	// Printed in reverse order
}

func main() {
	defer fmt.Println("world")


	defer_stack()


	fmt.Println("hello")
}
