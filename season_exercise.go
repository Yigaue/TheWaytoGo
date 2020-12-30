package main

import (
	"fmt"
	"time"
)
var month time.Month = time.Now().Month()
var season string
func  main() {
	fmt.Println(month)
	switch month {
	case 1, 2, 3:
	season = "Festive"
	case 4, 5, 6, 7:
	season = "Dry Season or Hamathan"
	default:
		season = "Rainy Season"
	} 
}