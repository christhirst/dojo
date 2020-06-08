package main

import (
	"fmt"
)

func main() {
	// var name type = expression
	// type or expression can be  omitted
	// zero value mech
	var s1 string
	fmt.Println(s1) // ""

	var i, j, k int // int int int
	fmt.Println(i, j, k)
	i, j, k = 1, 2, 3
	fmt.Println(i, j, k)

	var b, f, s2 = true, 2.3, "four" // bool float64, string
	fmt.Println(b, f, s2)

	// short variable declaration --->  name := expression
	// = is asignment
	// := deklaration
	fmt.Println(meineFn(s2))

	fmt.Println("Swapping Value ...")
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
}

func bla() string {
	return "bla"
}

func meineFn(s string) string {
	t := bla()
	return t
}
