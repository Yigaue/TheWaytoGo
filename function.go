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

func calculatePropertiesOfCircle(radius, pi float64) (area, perimeter float64) {
	perimeter = 2 * pi * radius
	area = pi * radius * radius

	return
}

func main() {
	price, no := 9000, 5
	result := calculateBill(price, no)
	fmt.Println(result)
	fmt.Println(calculateBill(7900, 6))

	perimeter, area := calculatePropertiesOfRectangle(6, 4)
	fmt.Println(perimeter, area)

	/** Name Return Value

	**/

	areaOfCircle, perimeterOfCircle := calculatePropertiesOfCircle(5, 3.143)

	fmt.Println(areaOfCircle, perimeterOfCircle)


}
