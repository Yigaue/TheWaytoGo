package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName string
	age int
	salary int
}

func main() {
	staff4 := Employee{
		firstName: "James",
		lastName: "Bond",
		salary: 8900,
		age: 43,
	}

	fmt.Println("First Name:", staff4.firstName)
	fmt.Println("Last Name:", staff4.lastName)
	fmt.Println("Age:", staff4.age)
	fmt.Println("Salary:", staff4.salary)
	staff4.salary = 6500000
	fmt.Println("New Salary:", staff4.salary)
}