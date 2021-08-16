package main

import (
	"fmt"
)

var height, weight, age int = 151, 68, 45

var (
	firstName,
	lastName string
	dob int
)

func main() {
	fmt.Println("height, weight, age:", height, weight, age)
	fmt.Println(firstName, lastName, dob)
	dob = 1998
	firstName = "John"
	lastName = "Doe"
	fmt.Printf("My name is %v %v, I was born in %d\n", firstName, lastName, dob)
}
