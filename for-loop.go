package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting loop")
	for i := 1; i <= 10; i ++ {
		if i == 7 {
			break
		}

		if i == 4 {
			continue
		}

		fmt.Printf("%d\n", i)
	}
}
