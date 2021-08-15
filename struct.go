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

	staff1 := Employee {
		firstName: "John",
		lastName: "Doe",
		age: 39, //The items define can be less than what is the struct declaration
				// e.g Salary is not included here. Duplicate field name not allowed.
				// the order of implementing the struct does not matter.
	}

	staff2 := Employee {
		firstName : "Jane",
		salary: 9000,
		age: 36,
		// firstName: "Doe", // not allow because firstName already exist
		// height: 56,// It won't fly because the struct employee does not have height in it.
	}

	fmt.Println("staf1:", staff1)
	fmt.Println("staff2:", staff2)
}