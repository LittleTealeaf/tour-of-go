package main

import "fmt"

const PI = 3.14

func main() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
