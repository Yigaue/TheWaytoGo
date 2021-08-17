package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	return price * no
}

func calculatePropertiesOfRectangle(length, width float64) (float64, float64) {
	perimeter := 2 * (length + width)
	area := length * width

	return perimeter, area
}

func main() {
	price, no := 9000, 5
	result := calculateBill(price, no)
	fmt.Println(result)
	fmt.Println(calculateBill(7900, 6))

	perimeter, area := calculatePropertiesOfRectangle(6, 4)
	fmt.Println(perimeter, area)
}
