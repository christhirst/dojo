package main

import (
	"fmt"
	"math"
)

const precission = 10000000 //000 //0000000

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < precission; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
