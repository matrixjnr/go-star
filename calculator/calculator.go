package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(CalculateFloat(1, 2, "+"))
	fmt.Println(CalculateFloat(1, 2, "-"))
	fmt.Println(CalculateFloat(1, 2, "*"))
	fmt.Println(CalculateFloat(1, 2, "/"))
}