package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a Vertex implements Abser

	// In the following line, v is a vertex, not *Vertex
	// a = v (this will not compile)

	fmt.Println(a.Abs())

	// Interfaces are implemented implicitly, meaning that as long as the methods are satisfied, it can be treated as such

	nil_underlying_values()
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func nil_underlying_values() {
	var i I

	var t *T
	i = t

	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	// Calling a value on a nil interface type will throw an error

	// also, interface {} applies to ALL values
}

func describe_all(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
