package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {println(a)}
func m() {
	a := "O"
	println(a)
}


// G, O, G