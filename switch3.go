package main

import (
	"fmt"
)

func main() {

	finger := 7
//	fmt.Printf("Finger %d is ", finger)

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger supplied")
	}

	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
		case "a", "e", "i", "o", "u" :
			fmt.Println("vowel") 
		default: 
			fmt.Println("Consonant")
	}

	num := 67
	switch {
		case num < 50:
			fmt.Println("Number is less than 50")
		case num > 50 && num < 100:
			fmt.Println("Number is greater than 50 but less than 100")
	}
}
