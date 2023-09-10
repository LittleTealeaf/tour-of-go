package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// [n]type
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// []T is a slice of elements T
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are references to arrays, they do not store data
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[0] = "XXX"
	fmt.Println(x, y)
	fmt.Println(names)

	// Array and slice literals
	q := []int{2, 3, 5, 7, 9, 11}
	fmt.Println(q)

	r := []bool{true, false, true, true, true, true}
	fmt.Println(r)

	z := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(z)

	// Default slice is the entire array

	slice_z := z[:]
	fmt.Println(slice_z)

	printSlice(primes[:])
	printSlice(primes[:0])
	printSlice(primes[:4])
	fmt.Println("When the start of the slice is not the first (is 2):")
	printSlice(primes[2:])

	// Nil Slices
	fmt.Println("This is a nil slice")
	var nil_slice []int
	printSlice(nil_slice)
	if nil_slice == nil {
		fmt.Println("Nil Slice is indeed a nil slice")
	}

	fmt.Println("Lets make some with MAKE")
	with_make()

	fmt.Println("Slices of Slices!")
	slices_with_slices()

	fmt.Println("Appending Slices")
	appending_slices()

	fmt.Println("Iteration!")
	iterating_over_arrays()
}

func with_make() {
	a := make([]int, 5)
	fmt.Println(a)

	b := make([]int, 0, 5)
	fmt.Println(b)

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b) = 4, cap(b) = 4
}

func slices_with_slices() {

	board := [][]string{
		// Alternatively string[]{"_", "_", "_"}
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appending_slices() {

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func iterating_over_arrays() {

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		// i is index
		// v is value
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d, ", value)
	}
	fmt.Println()
}
