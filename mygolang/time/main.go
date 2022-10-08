package main

import (
	"time"
	"fmt"
)

func main() {
	var myTime = time.Now()

	fmt.Println(myTime)

	fmt.Println(myTime.Format("01-03-2022"))
}