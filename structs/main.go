package main

import "fmt"

func main() {

	fmt.Println("These are structs")

	Sinmi := User{"Sinmi", "emmasimons141@gmail.com", true, 16}
	fmt.Println(Sinmi)
	fmt.Printf("Sinmi's details are: %+v\n", Sinmi)
	fmt.Printf("Name is %v and email is %v.", Sinmi.Name, Sinmi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
