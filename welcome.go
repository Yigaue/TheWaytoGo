package main

import (
	"fmt"
	"os"
	"./trans"
)

var twoPi = 2 * trans.Pi

func main() {
	greet := 3 > 1

	const eight string = "8"

	const SEVEN int = 7

	var (
		f int
		g   int = 3
		num int = 20
	)

	num = 8

	f = g

	var (
		a = os.Getenv("HOME")
		b = os.Getenv("USER")
		c = os.Getenv("GOROOT")
		d = os.Getenv("GOOS")
	)

	type Color int

	const (
		RED    Color = iota // 0
		ORANGE              // 1
		YELLOW              // 2
		GREEN               // ..
		BLUE
		INDIGO
		VIOLET // 6
	)

	const LOW = iota + 50

	if greet {
		fmt.Println(a, b, c, d, LOW, SEVEN, num, eight, f, g)
	}

	fmt.Printf("2*Pi = %g\n", twoPi)

	boo()

	fmt.Printf("something more")

	// for i := 0; i < 100; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	foo()
}

func boo() {
	var g int64 = 9223372036854775807
	fmt.Println("I am in boo", g)
}

func foo() {
	fmt.Println("Program exited")
}
