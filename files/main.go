package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "This needs to go in a file - Facebook.com"

	file, err := os.Create("./mylearning.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylearning.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)

	checkNilError(err)

	fmt.Println("Text data in file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
