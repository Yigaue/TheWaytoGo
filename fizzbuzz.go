package main

import (
	"fmt"
	// "math"
)

func main() {
	for i := 1; i <= 100; i++ {
		// fmt.Println(i)
		switch i {
		case modThreeORFive(i):
			fmt.Println("FizzBuzz")
		case modThree(i):
			fmt.Println("Fizz")
		case modeFive(i):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func modThree(n int ) int {
	if n % 3 == 0 {
		return n
	}
	return 0
}

func modeFive(n int) int {
	if n % 5 == 0 {
		return n
	}
	return 0
}

func modThreeORFive(n int) int {
	if n % 3 == 0 && n % 5 == 0 {
		return n
	}

	return 0
}
