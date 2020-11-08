/* moretypes13.go */
package main

import "fmt"

func ps(name string, s []int) {
	fmt.Printf("%s --> len=%d cap=%d --> %v\n", name, len(s), cap(s), s)
}

func main() {
	a := make([]int, 5)
	ps("a", a)
	b := make([]int, 0, 5)
	ps("b", b)
	_ = append(b, 1)
	c := b[0:2]
	ps("c", c)
	d := c[2:5]
	ps("d", d)
}
