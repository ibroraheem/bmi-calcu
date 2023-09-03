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
	if result < 18.0 {
		fmt.Printf("Your BMI is: %.2f. You are underweight", result)
		return result
	} else if result >= 18.0 && result < 25.0 {
		fmt.Printf("Your BMI is: %.2f. You have healthy weight", result)
		return result
	} else if result >= 25.0 && result < 30 {
		fmt.Printf("Your BMI is: %.2f. You are overweight", result)
		return result
	} else {
		fmt.Printf("Your BMI is: %.2f. You have obesity", result)
		return result
	}
}
