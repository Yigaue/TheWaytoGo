package main

import "fmt"

import "runtime"

func main () {

	var checkNumber = 3 > 1

	var str = ""

	if checkNumber {
		fmt.Printf("The value is true\n")
	}else {
		fmt.Printf("The value is false\n")
	}

	if str == "" {
		fmt.Printf("string is empty \n")
	} else {
		fmt.Printf("String is not empty \n")
	}

	println(runtime.GOOS)
	
	// var prompt = "Enter a digit, e.g 3"  + "or %s to quit. "

	// func init() {
	// 	if os  == "windows" {
	// 		prompt  = fmt.Sprintf(promt, "Ctrl + Z, Enter")

	// 	} else {
	// 		prompt = fmt.Sprintf(prompt, "Ctrl + D")
	// 	}
	// }
}