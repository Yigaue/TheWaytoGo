package main

import "fmt"

var char string = "G"

func main() {
	for i := 0; i < 7; i++ {
		// for j := 0; j < i; j++ {
		// 	fmt.Println(char)
		// }
		fmt.Println(char)
		char += char
	}
}