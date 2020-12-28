package main

import (
	"fmt"
	"strings"
)

var here = "HERE"
func main () {
	here = "I'm here, He is here"
	fmt.Printf("Here in string \"%s\" \"here\" was replaced  with \"there\" to become ", here)
	fmt.Printf("%s \n", strings.Replace(here, "here", "there", -1))
}