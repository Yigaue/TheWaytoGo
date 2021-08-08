package main

import (
	"fmt"
)

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func ax() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func bx() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	ax()
}

func main() {
	bx()
}
