package main

import (
	"fmt"
	"math"
)

const a = 5

const (
	d int = 3
	c int = 4
)

func main() {
	fmt.Println(a, c, d)
	f := math.Sqrt(float64(c)) // This works as against a const
	// const e = math.Sqrt(float64(c)) // This won't work.
	/*The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.*/
	fmt.Println(f)

	const n = "Sam"
	var name = n
	fmt.Printf("type %T value %v\n", name, name)
}
