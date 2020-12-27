package main

import (
	"fmt"
	"strings"
)

func main () {
	var str string  = "Go is a nice programming language"

	fmt.Printf("T/F ? Does the string \"%s\" has prefix %s ? ", str, "Go")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Go"))
}