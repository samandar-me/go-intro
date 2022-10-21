package main

import "fmt"

func main() {
	var user = User{"Samandar", "Asiydinov", "samandar.go@gmail.com", 15, true}
	fmt.Println(user)
	fmt.Printf("Samandar details are: %+v\n", user)
	fmt.Printf("Name is %v", user.Name)
}

type User struct {
	Name string
	
	LastName string
	Email string
	Age int
	Gender bool
}