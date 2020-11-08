/* moretypes22.go */

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	a, err := m["Answer"]
	fmt.Println(a)
	fmt.Println(err)
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
