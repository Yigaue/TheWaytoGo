package main

import (
	"fmt"
)

func main() {
	var a [3]int // int array a with length 3
	a[0] = 12
	a[1] = 78
	a[2] = 50
	b := [10]int{12, 50}      // the compiler fils the vacant space the zero value of the type.
	c := [...]int{14, 97, 43} // the compiler determines the length
	fmt.Println(a, b, c)

	countries := [...]string{"USA", "Germany", "Israel", "China"}
	fmt.Println(countries)
	fmt.Printf("The countries are %v\n", countries)

	for i := 0; i < len(countries); i++ {
		fmt.Printf("The %v country is %v\n", i+1, countries[i])
	}

	k := [...]float64{89.9, 73.3, 34.0, 89.4}
	sum := float64(0)

	for i, v := range k {
		fmt.Printf("the %d element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Printf("\nSum of all elements of a %.2f\n", sum)

}
