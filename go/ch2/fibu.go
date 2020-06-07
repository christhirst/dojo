package main

import (
	"flag"
	"fmt"
)

var iterations = flag.Int("iterations", 1, "Counts of iterations")

func main() {
	flag.Parse()
	fmt.Println(fib(*iterations))
}

func fib(n int) int {
	var i, x, y = 0, 0, 1
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
