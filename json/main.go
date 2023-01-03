package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"price"`
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println(("Welcome to JSON video"))
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 290, "LearcodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"ReactJs Bootcamp", 210, "LearcodeOnline.in", "adc123", []string{"mobile-dev", "js"}},
		{"Laravel Bootcamp", 240, "LearcodeOnline.in", "ebc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"price": 210,
		"Platform": "LearcodeOnline.in",
		"tags": ["mobile-dev","js"]
	}
	`)

	var lcoCourses course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourses)
}
