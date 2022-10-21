package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	clubs := []string{"Tottenham", "Real Madrid", "Barcelona", "Bavariya", "Man City", "Man UTD", "Psg", "Juventus"}
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(8) + 1
	fmt.Println("Random club is ", clubs[randomNumber])



	switch randomNumber {
	case 1:
		// todo
	case 2:
		// todo
	} 

}