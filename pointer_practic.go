package main 

import (
	"fmt"
	"math"
)

func main () {
	const NUM = 5
	const DIGIT  = math.Pi * NUM // No error is thrown if the CONST is not use, but an error will be thrown if the variable is declared but not used
	var digit = DIGIT

	fmt.Println(digit)
	fmt.Println(&digit)
	fmt.Println(*&digit)
}