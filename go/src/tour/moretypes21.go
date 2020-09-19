/* moretypes21.go */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//Top Level is just a type name, Vertex can be omited
var m = map[string]Vertex{
	"Bell Labs": {30, -30},
	"Google":    {20, -20},
}

func main() {
	fmt.Println(m)
}
