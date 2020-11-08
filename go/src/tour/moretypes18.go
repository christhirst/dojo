/* moretypes18.go */

package main

import (
	"golang.org/x/tour/pic"
)

func calc(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8

	for y := 0; y < dy; y++ {
		ret = append(ret, make([]uint8, 0))
		for x := 0; x < dx; x++ {
			ret[y] = append(ret[y], calc(x, y))
		}

	}
	return ret

}

func main() {
	pic.Show(Pic)
}
