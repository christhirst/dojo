package main

import (
	"fmt"
	"strings"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	hashmap := make(map[string][]string)
	hashmap["mymedals"] = medals
	fmt.Println(strings.Join(medals, " "))
	fmt.Println(hashmap["mymedals"])
}
