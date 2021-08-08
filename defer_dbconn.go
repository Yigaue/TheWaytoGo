package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("Ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("Ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Deferring the database disconnect.")
	defer disconnectFromDB() //function call here with defer
	fmt.Println("Doing some DB operations...")
	fmt.Println("Oops! some crash or network error...")
	fmt.Println("Returning from function here!")
	return // terminate the program
	//deffered function executed here just before actually returning, even if there is 
	//return or abnormal termination before
}