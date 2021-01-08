package main

import "fmt"

var min, max int
func main() {
	min, max  = MinMax(78, 56)
	fmt.Printf("Minimum is: %d, maximum is: %d\n", min, max)
}

func MinMax(a int, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}

	return
}