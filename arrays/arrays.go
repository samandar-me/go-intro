package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	var clubList [5]string

	clubList[0] = "Tottenham"
	clubList[1] = "Real Madrid"
	clubList[2] = "Man City"
	clubList[3] = "Bavariya"
	clubList[4] = "Psg"

	clubList[len(clubList) / 2] = "Juventus"

	// fmt.Println("ClubList , ", clubList)
	// fmt.Println("Lenght of clublist ",len(clubList))

	
}