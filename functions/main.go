package main

import "fmt"

func main() {

	fmt.Println("Enter two nubmers: ")
	
	var a int
	_, e := fmt.Scanf("%2d", a)

	var b int
	_, e2 := fmt.Scanf("%2d", b)

	menu()
	var option int
	fmt.Println("Enter option:")
	_, err := fmt.Scanf("%d", &option)

	if err != nil {
		fmt.Println(err)
	}

	if e != nil {
		fmt.Println(e)
	}

	if e2 != nil {
		fmt.Println(e2)
	}

	switch option {
	case 1:
		print(add(a, b))
	case 2:
		print(mul(a,b))
	case 3:
		print(sub(a,b))
	case 4:
		print(div(a,b))
	default:
		fmt.Println("Invalid option!")
	} 

	middle, msg := middleValue(23,23,423,323)



	fmt.Println(middle)
	fmt.Println(msg)
}

func print(msg int) {
	fmt.Println(msg)
}

func menu() {
	fmt.Println("1. Add")
	fmt.Println("2. Mul")
	fmt.Println("3. Sub")
	fmt.Println("4. Div")
}

func add(a int, b int) int {
	return a + b
}

func div(a int , b int) int {
	if a > b {
		return a / b
	}
	return b / a
}

func mul(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}

func middleValue(values ...int) (int , string) {
	var sum = 0

	for val := range values {
		sum += val
	}
	return sum / len(values), "Hi"
}