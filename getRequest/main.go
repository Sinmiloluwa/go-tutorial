package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
