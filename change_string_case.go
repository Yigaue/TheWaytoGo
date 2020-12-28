package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string
	fmt.Printf("The original string is %s\n", orig)

	lower = strings.ToLower(orig)
	fmt.Printf("The lower case string is %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The upper case string is %s\n", upper)
}