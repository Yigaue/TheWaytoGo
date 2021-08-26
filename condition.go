package main

import (
	"fmt"
)


func main() {
	num := 300

	if num % 2 == 0 {
		fmt.Println("number is even:", num)
		return
	}

	fmt.Println("number is odd:", num)
}