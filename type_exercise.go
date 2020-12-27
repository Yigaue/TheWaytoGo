package main

import (
	"fmt"
)

type Rope string

func main() {
	var name Rope
	name = "Gospel"
	fmt.Printf("My name is %s \n", name)
}