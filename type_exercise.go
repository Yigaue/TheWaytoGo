package main

import (
	"fmt"
)

type Rope string

func main() {
	var name Rope
	name = "Gospel"
	// t := &name[5]
	fmt.Println(&name)
}