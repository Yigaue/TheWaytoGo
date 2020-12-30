package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "123"

	an, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("The string %s is not an integer \n", str)
		return 
	}

	fmt.Printf("The integer is %d\n", an)
}