package main

import (
	"fmt"
	"strings"
)

func main () {
	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)

	fmt.Printf("Splitted in slice: %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}

	str2 := "Go1|The ABC of Go|25"
	s12 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", s12)
	for _, val := range s12 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()

	str3 := strings.Join(s12, ";")
	fmt.Printf("s12 joined ;: %s\n", str3)
}