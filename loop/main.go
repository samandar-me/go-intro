package main

import (
	"fmt"
)

func main()  {
	clubs := []string{"Tottenham", "Real Madrid", "Barcelona", "Bavariya", "Man City", "Man UTD", "Psg", "Juventus"}

	for k := 0; k < len(clubs); k++ {
		fmt.Println(clubs[k])
	}
	for i := range clubs {
		fmt.Println(clubs[i])
	}
	for index, value := range clubs {
		fmt.Printf("Index %v, value %v\n", index, value)
	}
	var counter int = 1

	for counter < 50 {
		fmt.Println(counter)
		counter++
	}
}