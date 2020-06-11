package main

import (
	"fmt"
)

/*
& bitwise AND
| bitwise OR
^ bitwise XOR or complement
&^ bitclear ( AND NOT )
<< left shift
>> right shift
*/
func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1 << 1
	x = ^((2 << x) >> 1)
	fmt.Printf("%08b\n", x)

	var y uint8 = 1<<1 | 1<<5
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&^y)
	fmt.Printf("%08b\n", x^y)
}
