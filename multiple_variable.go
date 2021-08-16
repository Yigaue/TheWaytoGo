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

var (
	jobTitle string = "Software engineer"
	gender = "male"
	is_married = true
	nofKids int = 2
)

func main() {
	fmt.Println("height, weight, age:", height, weight, age)
	fmt.Println(firstName, lastName, dob)
	dob = 1998
	firstName = "John"
	lastName = "Doe"
	fmt.Printf("My name is %v %v, I was born in %d\n", firstName, lastName, dob)
	fmt.Printf("I am a %v. My gender is %v, married(%v) with %v kids\n", jobTitle, gender, is_married, nofKids)

	firstChild, childAge := "Sam", 16

	fmt.Printf("My first child is %v %v and he is %d\n", firstChild, lastName, childAge)
}
