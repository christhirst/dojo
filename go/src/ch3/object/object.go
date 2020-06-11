package object

import (
	"fmt"
)

//MyObject bla
type MyObject int64

//GetValue MyObject function
func (m MyObject) GetValue() string {
	return fmt.Sprintf("%d", m)
}
