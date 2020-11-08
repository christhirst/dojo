package main

import "fmt"

func ps(s []int) {
	fmt.Printf("len=%d, cap=%d --> %v\n", len(s), cap(s), s)
}

func main() {

	s := []int{2, 3, 5, 7, 11, 13}
	ps(s)
	s = s[:0]
	ps(s)
	s = s[:cap(s)]
	ps(s)
	s = s[:4]
	ps(s)
	//schmeisst die 1ten 2 Werte wirklich weg, cap wird kleiner ...
	s = s[2:]
	ps(s)
}
