package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	/* p  = &Vertex{1, 2} */
	p = v3
)

func main() {
	p.X = 4
	fmt.Println(p.X)
	fmt.Println(v3.X)
}
