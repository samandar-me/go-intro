package main

import "fmt"

func main() {

	var ptr2 *int 

	myNumber := 18

	var ptr = &myNumber

	fmt.Print(ptr)
	fmt.Println(*ptr)

	*ptr += 2
	fmt.Println(*ptr)
	fmt.Println(myNumber)
	fmt.Println(ptr2)
}