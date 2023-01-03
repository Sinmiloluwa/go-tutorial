package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("world")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
