package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=300"

func main() {
	fmt.Println("Welcome to handling URL")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)

	qparames := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparames)
	fmt.Println(qparames["coursename"])

	for _, val := range qparames {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutorials",
		RawPath: "user=sinmi",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
