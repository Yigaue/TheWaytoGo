package main

import (
	"fmt"
	"strings"
)

func main () {
	var str = " A string with leading and trailling spaces "
	var str1 = "\t A string with tab spaces and carraige return \n"
	fmt.Printf("A string with leading and trailling spaces trimed: ")
	fmt.Printf(" \"%s\" \n", strings.TrimSpace(str))
	fmt.Printf("A string with tabs and next spaces trimmed:")
	fmt.Printf("\"%s\"\n", strings.Trim(str1, "\n\t"))	
}