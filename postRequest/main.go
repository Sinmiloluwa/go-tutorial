package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	requestBody := strings.NewReader(`
		{
			"coursename": "Learn golang",
			"price":"free",
			"platform":"simonslearn.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
