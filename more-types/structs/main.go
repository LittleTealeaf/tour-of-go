package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5

	fmt.Println(v)
	
	p := &v

	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p2 = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, p2)


}
