package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Apple", "Tomato", "Orange"}
	fmt.Printf("there you have it %T\n", fruits)

	fruits = append(fruits, "Banana", "Watermelon")
	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 70
	highScores[1] = 80
	highScores[2] = 85
	highScores[3] = 90

	highScores = append(highScores, 92, 45, 63)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slice based on the index

	var courses = []string{"reactjs", "javascript", "ruby", "php"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
