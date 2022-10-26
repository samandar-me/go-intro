package main

import "fmt"

func main() {
	defer fmt.Println("Golang")
	defer fmt.Println("Kotlin")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i := 10; i >= 0; i-- {
		defer fmt.Println(i)
	}
	
}