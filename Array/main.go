package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruit [4]string

	fruit[0] = "garetd"
	fruit[1] = "Peach"
	fruit[2] = "Tomato"

	fmt.Println("Fruits are: ", fruit)
}
