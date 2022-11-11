package main

import "fmt"

const grain = "Money"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int

	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pointer is: ", *ptr)
}
