package main

import (
	"fmt"
	"strings"
)

func main () {
	var str string = "God is good, all the time"
	fmt.Printf("T/F ? Does the string \"%s\" contains %s ? ", str, "good")
	fmt.Printf("%t\n", strings.Contains(str, "good"))
}