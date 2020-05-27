package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		var start = time.Now()
		fmt.Println(os.Args[1:])
		fmt.Println(time.Now().Sub(start))
		start = time.Now()
		fmt.Println(strings.Join(os.Args[1:], " "))
		fmt.Println(time.Now().Sub(start))
	}
}
