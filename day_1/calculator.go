package main

import (
	"fmt"
)

func main() {
	var first, second, option int
	for {
		fmt.Println("Enter your first number")
		fmt.Scanln(&first)

		fmt.Println("Enter your second number")
		fmt.Scanln(&second)

		fmt.Printf("Enter option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Addition: ", add(first, second))
		case 2:
			fmt.Println("Subtraction: ", sub(first, second))
		case 3:
			fmt.Println("Multiplications: ", mul(first, second))
		case 4:
			fmt.Println("Division: ", div(first, second))
		default:
			fmt.Println("No option is selected...")
		}
	}
}

func add(first int, second int) int {
	return first + second
}

func sub(first int, second int) int {
	return first - second
}

func mul(first int, second int) int {
	return first * second
}

func div(first int, second int) int {
	return first / second
}
