package main

import "fmt"

func main() {
	countries := make(map[string]string)

	countries["ENG"] = "England"
	countries["NIG"] = "Nigeria"
	countries["ARG"] = "Argentina"

	fmt.Println("List of countries: ", countries)
	fmt.Println("ENG short for ", countries["ENG"])

	delete(countries, "ARG")
	fmt.Println("Countries: ", countries)

	for _, value := range countries {
		fmt.Printf("For key , value is %v\n", value)
	}
}
