package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", Multiply3Nums(2, 5, 6))
	// var i1 int = Multiply3Nums(2, 5, 6)
	// fmt.Printf("Multiply 2 * 5 * 6 = %d\n", i1)
}

func Multiply3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}