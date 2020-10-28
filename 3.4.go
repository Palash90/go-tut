// Automatic type recognition

package main

import (
	"fmt"
)

func main() {

	x := 1
	y := 2
	avg := (x + y) / 2

	// Something like python tuples
	x1, x2 := 1.0, 2.0
	avg1 := (x1 + x2) / 2

	fmt.Printf("avg=%v, type=%T\n", avg, avg)
	fmt.Printf("avg1=%v, type=%T\n", avg1, avg1)
}
