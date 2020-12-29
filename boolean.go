package main

import "fmt"

import "runtime"

var prompt = "Enter a digit, e.g 3"  + "or %s to quit. "

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

	 var abs = Abs(-8)
	 fmt.Println(abs)

	 println(isGreater(2, 5))
}

func init() {
		if runtime.GOOS  == "windows" {
			prompt = fmt.Sprintf(prompt, "Ctrl + Z, Enter")

		} else {
			prompt = fmt.Sprintf(prompt, "Ctrl + D")
		}
	}

func Abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if (x > y) {
		return true
	}

	return false
}
