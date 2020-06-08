package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Print(sep)
	for _, val := range flag.Args() {
		fmt.Println(val)
	}
	newFN()
}

func newFN() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func newInt() *int { /*
		var dummy int
		return &dummy */
	return new(int)
}
