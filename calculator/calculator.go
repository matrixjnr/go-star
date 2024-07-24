package main

import (
	"fmt"
)

// a calculator that follows bodmas
func Calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Invalid Operator")
		return 0
	}
}

// a calculator that follows bodmas
func CalculateFloat(a float64, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Invalid Operator")
		return 0
	}
}

// a calculator that follows bodmas
func CalculateIntegers(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Invalid Operator")
		return 0
	}
}

func main() {
	fmt.Println(Calculate(1, 2, "+"))
	fmt.Println(Calculate(1, 2, "-"))
	fmt.Println(Calculate(1, 2, "*"))
	fmt.Println(CalculateFloat(1, 2, "/"))
}