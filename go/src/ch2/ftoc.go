package main

import "fmt"

func main() {
	//:= declaration
	//= asignment
	a := 0
	var b = 0
	var c int
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Freezing: %g, %g", freezingF, ftoc(freezingF))
	fmt.Printf("Boiling: %g, %g", boilingF, ftoc(boilingF))
	fmt.Println(a, b, c)
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
