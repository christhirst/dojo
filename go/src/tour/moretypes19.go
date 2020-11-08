/* moretypes19.go */

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.0, -40.0}
	fmt.Println(m["Bell Labs"])
}
