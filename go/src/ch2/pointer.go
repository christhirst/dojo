package main

import (
	"fmt"
)

func main() {
	/* x := 1
	p := &x

	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)

	meinFN()
	fmt.Println(p)
	*/
	p3 := f()
	fmt.Println(p3)
	fmt.Println(p2)
	fmt.Println(&p2 == &p3)
	var a int
	var b int
	meinPlus(&a)
	meinPlus(&b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b)
	meinPlus(&b)
	fmt.Println(b)
}

func meinPlus(val *int) {
	*val++
}
func meinFN() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}

var p2 = f()

func f() *int {
	v := 1
	fmt.Println(v)
	return &v
}
