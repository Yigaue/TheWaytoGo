package main

import "fmt"

func main () {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	var intP *int // intP is a variable(pointer) that stores memory location of value(variable) of type int. Here it is declared but not assigned. The "*" indicates that intP is a pointer variable and not and ordinary variable. This means intP will store a pointer of value of type int.

	intP = &i1 // Here we assigned the memory location of variable i1 to intP. That makes intP a pointer.
	fmt.Printf("The value of memory location %p is %d\n", intP, *intP)
}