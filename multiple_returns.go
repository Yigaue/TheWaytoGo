package main

import (
	"fmt"
)

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	printValues()
	numx2, numx3 = getX2AndX3_2(num)
	printValues()
}

func printValues() {
	fmt.Printf("num = %d, 2X num = %d, 3X num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}