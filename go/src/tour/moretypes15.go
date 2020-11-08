/* moretypes15.go */

package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d --> %v\n", len(s), cap(s), s)
}

func extendSlice(count int, s []int) []int {
	for i := 0; i < count; i++ {
		s = append(s, i)
		fmt.Printf("cap: %d\n", cap(s))
	}
	return s
}

func main() {
	var count = 1000
	var s []int

	s = extendSlice(count, s)
	printSlice(s)
}
