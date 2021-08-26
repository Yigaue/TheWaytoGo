package main

import ("fmt")

func main() {
	if num <= 50 {
		fmt.Println("Number is less than 50:", num)
	} else if num <= 100 && num >= 50 {
		fmt.Println("Number is between 50 and 100:", num)
	} else {
		fmt.Println("number is greater than  100:", num)
	}
}