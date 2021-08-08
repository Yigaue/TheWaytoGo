package main

import (
	"fmt"
)

type Rope string

func main() {
	var name Rope
	name = "Messiah"
	// t := &name[5]
	fmt.Println(&name)
}