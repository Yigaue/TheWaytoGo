package main

import (
	"fmt"
	"strings"
)

func main () {
	var str = "A string with spaces between words"
	fmt.Printf("%s\n", strings.Fields(str)) // on whitespace
	str  = "A, string, with, comma, and spaces between words"
	fmt.Printf("%s\n", strings.Split(str, ","))
}