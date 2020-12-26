package main

import (
	"math"
	"fmt"
)

func main() {
	result, err := Uint8FromInt(254)
	fmt.Println(result, err)
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}

	return 0 , fmt.Errorf("%d is out of the unint8 range", n)
}