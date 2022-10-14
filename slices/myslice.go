package main

import (
	"fmt"
	"sort"
)

func main() {
	var languageList = []string{"Kotlin", "Java", "Golang", "Swift", "C", "C++", "Python"}
	fmt.Printf("Type of languagelist is %T\n", languageList)

	languageList = append(languageList, "JavaScript", "Dart")

	fmt.Println(languageList)
	languageList = append(languageList[:5])
	fmt.Println(languageList)

	ints := make([]int, 5)
	ints[0] = 6
	ints[1] = 1
	ints[2] = 5
	ints[3] = 8
	ints[4] = 2

	fmt.Println(ints)

	ints = append(ints, 9,0,3)
	fmt.Println(ints)

	sort.Ints(ints)
	fmt.Println(ints)

	var clubList = []string{"Tottenham", "Real Madrid", "Barcelona", "Bavariya", "Man City", "Man UTD", "Psg", "Juventus"}
	var index = 2

	fmt.Println(clubList)

	clubList = append(clubList[:index], clubList[index+1:]...)
	fmt.Println(clubList)
}