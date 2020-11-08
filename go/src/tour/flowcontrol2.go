package main

import "fmt"

func main() {
	var sum = 0
	for sum < 1000 {
		sum += 1000
	}

	fmt.Println(sum)
}
