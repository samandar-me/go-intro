package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Enter your sim code")
	_, err := fmt.Scanf("%d", &input)

	if err != nil {
		fmt.Println(err)
	}
	var result string
	
	if input == 98 || input == 99 {
		result = "Uzmobile"
	} else if input == 94 || input == 93 {
		result = "Ucell"
	} else if input == 90 || input == 91 {
		result = "Beeline"
	} else if input == 33 {
		result = "Humans"
	} 
		fmt.Println("You have using ", result)

		if num := 4; num < 10 {
			fmt.Println(num)
		} else {
			fmt.Println(num % 2)
		}
}