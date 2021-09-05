package main

import ("fmt")
// Note: Slices - A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.
func main() {
	//A slice with elements of type T is represented by []T

	a := [5]int{3, 4, 6, 90, 76}
	var b[]int = a[1:4] //create a slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{5, 207, 54} // creates an array and return a slice reference

	fmt.Println(c)

	darray := [...]int{90, 50, 34, 94, 55, 42, 12}
	dslice := darray[2:5]
	fmt.Println("array before modification", darray)
	fmt.Println("slice before modication", dslice)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after modification", darray)
	fmt.Println("slice after modication", dslice)
}