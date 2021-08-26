package main 

import ("fmt")

func main() {
	if num := 11; num % 2 == 0 {
		fmt.Println(num, "is even")
		return
	} else {
		fmt.Println(num,"is odd")
	}
}