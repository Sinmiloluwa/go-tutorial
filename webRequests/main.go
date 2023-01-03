package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://theconnectedawards.com"

func main() {
	fmt.Println("WEB REQUEST")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
