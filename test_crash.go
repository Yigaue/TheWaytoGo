package main

func main() {
	var p *int = nil // A nil pointer dereference is illegal and
					// makes a program crash
	*p = 0
}