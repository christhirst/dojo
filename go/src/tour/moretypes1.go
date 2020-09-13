/* moretypes1.go */

package main

import "fmt"

func main() {
	/* var p *int
	i := 42
	p = &i

	fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(*p)
	*/

	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
