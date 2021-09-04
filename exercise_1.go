package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 42
		b = iota
		c = iota
	)

	fmt.Printf("%d\t%b\t%#x\t%T\n", a, a, a, a)
	fmt.Printf("%d\t%b\t%#x\t%T\n", b, b, b, b)
	fmt.Printf("%d\t%b\t%#x\t%T\n", c, c, c, c)
}
