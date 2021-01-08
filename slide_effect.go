package main

import "fmt"

// This functions changes reply:

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n // reply points to the mempry location of n
				// and *reply gives the value occupy by that memory location
	// println(*reply)
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50	
}