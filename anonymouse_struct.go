package main

import (
	"fmt"
)

/*---------------------------
Note: It is possible to declare struct without 
creating a new data type. This type of structs
are called anonymouse struct.
-----------------------------
*/

func main() {
	staff3 := struct {
		firstName string
		lastName string
		age int
		salary int
	}{
		"Mary", // you can leave the name field 
		//but you must follow the order define in the struct above
		"Joe",
		29,
		4000,
	}

	fmt.Println("staff 3:", staff3)
}