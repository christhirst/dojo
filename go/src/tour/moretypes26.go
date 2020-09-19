/* moretypes26.go */

package main

import "fmt"

func fibonacci() func() int {
	ret := []int{}
	fibo := func() int {
		if len(ret) == 0 {
			ret = append(ret, 0)
			return 0
		}
		if len(ret) == 1 {
			ret = append(ret, 1)
			return 1
		}
		ret = append(ret, ret[len(ret)-1]+ret[len(ret)-2])
		ret = ret[1:]
		fmt.Println(len(ret), cap(ret))
		return ret[len(ret)-1]
	}
	return fibo
}

func main() {
	f := fibonacci()
	for range make([]int, 10) {
		fmt.Println(f())
	}
	fmt.Println("Ende")
}
