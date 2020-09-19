/* moretypes23.go */

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, w := range strings.Fields(s) {
		ret[w]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
