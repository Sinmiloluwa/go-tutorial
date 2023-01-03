package main

import "fmt"

func main() {

	fmt.Println("These are structs")

	Sinmi := User{"Sinmi", "emmasimons141@gmail.com", true, 16}
	fmt.Println(Sinmi)
	fmt.Printf("Sinmi's details are: %+v\n", Sinmi)
	fmt.Printf("Name is %v and email is %v.\n", Sinmi.Name, Sinmi.Email)
	Sinmi.GetStatus()
	Sinmi.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "testdio@den.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
