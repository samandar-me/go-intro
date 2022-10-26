package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var fileName = "./myFile.txt"
	var content string = "I am learning golang development"
	file, err := os.Create(fileName)
	checkNil(err)
	length, err := io.WriteString(file, content)
	checkNil(err)
	fmt.Println("File length is ", length)
	defer file.Close()

	readFile(fileName)
}

func readFile(fileName string) {
	databyte,err := ioutil.ReadFile(fileName)
	checkNil(err)
	fmt.Println("File is ", string(databyte))
}
func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}