package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your favorite programming language: ")

	input, _:= reader.ReadString('\n')
	fmt.Println("Your favorite language is", input)
}