/* meins.go */

package main

import "fmt"

func main() {
	var a [10]int
	var b = a[:]
	var c = &a
	var d = make([]int, 10)
	fmt.Println(a)
	a[1] = 1
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(d)
	var e = []int{0}
	fmt.Println(e)
	e = append(e, 100)
	fmt.Println(e)
	fmt.Println(cap(e), len(e))
	e = append(e, 100)
	fmt.Println(cap(e), len(e))
}
