package main

import (
	"fmt"
	"time"
)

func main() {
	presentDate := time.Now()
	fmt.Println(presentDate)

	fmt.Println(presentDate.Format("01-02-2006 Monday"))

}
