package main

import (
	"fmt"
)

func main() {
	var num1 int = 100

	switch num1 {
	case 98, 88:
		fmt.Println("it's equal 98")

	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}