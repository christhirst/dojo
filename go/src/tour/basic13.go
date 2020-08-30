package main

import (
	"fmt"
	"math"
)

var i = 42
var f float64 = float64(i)
var u uint = uint(f)

func main() {
	var x, y int = 3, 4
	f = math.Sqrt(float64(x*y + x*y))
	z := f
	fmt.Println(x, y, z)
}
