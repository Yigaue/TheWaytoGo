package main

import "fmt"

var sum, product, difference int
var sum1, product1, difference1 int

func main() {
	sum, product, difference = basicArithmetic(2, 3)
	sum1, product1, difference1 = basicArithmetic_2(7, 3)
	fmt.Println(sum, product, difference)
	fmt.Println(sum1, product1, difference1)
}

func basicArithmetic(x1, x2 int) (int, int, int) {
	sum = x1 + x2
	product = x1 * x2
	difference = x1 - x2

	return sum, product, difference
}

func basicArithmetic_2(x1, x2 int) (sum, product, difference int) {
	sum = x1 + x2
	product = x1 * x2
	difference = x1 - x2

	return
}
