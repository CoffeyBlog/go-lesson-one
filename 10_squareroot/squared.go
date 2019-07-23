//Basic Square root program in Go

package main

import (
	"fmt"
	"math"
)

func main() {

	var i int32

	i = int32(math.Sqrt(15.0))

	fmt.Print("Value of i:", i)
}
