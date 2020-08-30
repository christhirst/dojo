package main

import (
	"fmt"
	"unicode/utf8"
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

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) //438 666 0666

	//var f float32 = 16777216 //true
	//var f float64 = 16777216 //false
	var f int32 = 16777216 //false
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1)

	const e = 2.71828
	/*
		g --> chooses the most compact representation
		e --> exponent representation
		f --> noexponent representation
	*/
	fmt.Printf("%[1]g %[1]e %[1]f", e)

	/*
		IEEE 754
		negative and positive infinities
		zero devisioned by zero --> NaN ( not a Number )
	*/
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	fmt.Println(compute())

	myString()
}

func compute() (float64, bool) {
	failed := true
	if failed {
		return 0, false
	}
	return 1.0, true
}

func myString() {
	/*
		len() returns bytes
		index operand returns the n-th byte, not rune
	*/
	s := "hello, worldä"
	fmt.Println(len(s))            //12
	fmt.Println(s[0], s[len(s)-1]) //104 164
	fmt.Println("\164")

	/*
		c := s[len(s)] // panic
		fmt.Println(c)
		... working with utf-8 needs 3 bytes per rune to work with
	*/
	fmt.Println(s[0:5])
	fmt.Println(s[:5])

	//strings are immutable
	fmt.Println(&s)
	t := s
	s = " "
	fmt.Println(&s)
	fmt.Println(t)
	fmt.Println(`\`)

	/*
		Cannot be assigned ...
		s[0] = 'L'
	*/

	s = "hellow world"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Println(r)
		i += size
	}

}
