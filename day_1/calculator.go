package main

import (
	"fmt"
	"math"
)

func main() {
	var first, second, option int
	for {
		fmt.Printf("Enter option: ")
		fmt.Scanln(&option)

		fmt.Println("Enter your first number")
		fmt.Scanln(&first)

		if option < 5 {
			fmt.Println("Enter your second number")
			fmt.Scanln(&second)
		}

		switch option {
		case 1:
			fmt.Println("Addition: ", add(first, second))
		case 2:
			fmt.Println("Subtraction: ", sub(first, second))
		case 3:
			fmt.Println("Multiplications: ", mul(first, second))
		case 4:
			fmt.Println("Division: ", div(first, second))
		case 5:
			fmt.Println("Sin: ", sin(float64(first)))
		case 6:
			fmt.Println("Cos: ", cos(float64(first)))
		case 7:
			fmt.Println("tan: ", tan(float64(first)))
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

func sin(degree float64) float64 {
	radian := degree * math.Pi / 180
	return math.Sin(radian)
}

func cos(degree float64) float64 {
	radian := degree * math.Pi / 180
	return math.Cos(radian)
}

func tan(degree float64) float64 {
	radian := degree * math.Pi / 180
	return math.Tan(radian)
}
