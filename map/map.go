package main

import "fmt"

func main()  {
	clubs := make(map[string]string)
	clubs["Tot"] = "Tottenham"
	clubs["Rma"] = "Real Madrid"
	clubs["FCB"] = "Barcelona"
	clubs["MCI"] = "Man City"
	clubs["MYU"] = "Man UTD"
	clubs["Juv"] = "Juventus"
	clubs["Bay"] = "Bayern Munich"
	clubs["PSG"] = "Paris Saint-Germain"

	fmt.Println("List of map", clubs)
fmt.Println(clubs["Tot"])
	delete(clubs, "MYU")
	fmt.Println("List of map", clubs)

	for key, value := range clubs {
		fmt.Printf("For key %v - valus is %v, ", key ,value)
	}
}