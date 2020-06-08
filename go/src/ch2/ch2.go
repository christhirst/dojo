package main

import (
	"ch2/tempconv0"
	"fmt"
	"os"
)

func x() int {
	return 0
}
func main() {
	fmt.Printf(tempconv0.AbsoluteZeroC.String())
	if a := x(); a == 0 {
		fmt.Println(a)
	} else {
		//
	}

	if f, err := os.Create("./myfile"); err != nil {
		if a, err := f.WriteString("bla"); err != nil {
			fmt.Println(a)
		} else {
			fmt.Println(a)
		}
	} else {
		fmt.Println("blub")
		fmt.Println(err)
	}
}
