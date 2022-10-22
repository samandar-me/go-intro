package main

import "fmt"

func main() {
	club1 := FootballClub{"Tottenham", "Apl", 3, "Tot", "Konte"}
	PrintClubInfo(club1)
}
type FootballClub struct {
	Name string
	Leagues string
	Place int
	ShortName string
	CoachName string
}

func PrintClubInfo(club FootballClub) {
	fmt.Println("Club name is ", club.Name)
	fmt.Println("Club league is ", club.Leagues)
	fmt.Println("Place ", club.Place)
	fmt.Println("Short Name", club.ShortName)
	fmt.Println("Coach Name is ", club.CoachName)
} 