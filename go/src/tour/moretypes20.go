/* moretypes20.go */

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.0, -40.0,
	},
	"Google": Vertex{
		50.0, -50,
	},
}

func main() {
	fmt.Println(m)
}
