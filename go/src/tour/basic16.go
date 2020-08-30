package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {

	return x
}

func main() {
	fmt.Println(needInt(1))
	fmt.Println(needFloat(0.1))
}
