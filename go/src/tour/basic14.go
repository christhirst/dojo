package main

import "fmt"

var i = 42

func main() {
	v := i
	f := 3.142
	g := 0.876 + 0.5i
	v2 := f
	v3 := g
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n", v2)
	fmt.Printf("v is of type %T\n", v3)
}
