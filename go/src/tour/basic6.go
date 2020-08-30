package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	fmt.Println(swap(1, 2))
}
