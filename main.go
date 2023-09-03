package main

import (
	"fmt"
	"math"
)

func main() {
	var weight, height float64
	calculate(weight, height)
}

func calculate(weight, height float64) float64 {
	fmt.Println("Welcome to my BMI calculator!!!")
	fmt.Println("Enter your weight (kg):")
	fmt.Scan(&weight)
	fmt.Println("Enter your height (m):")
	fmt.Scan(&height)
	denom := math.Pow(height, 2)
	result := weight / denom
	fmt.Printf("Your BMI is: %.2f", result)
	return result
}
