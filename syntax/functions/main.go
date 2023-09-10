package main

import "fmt"


func add_full(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func multiple_results(x, y string) (string, string) {
	return y, x
}

func named_returns(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add_full(10, 13))
	fmt.Println(add(42, 13))

	a, b := multiple_results("hello", "world");

	fmt.Println(a, b)


	fmt.Println(named_returns(17))
}
