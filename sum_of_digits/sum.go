package main

import "fmt"

// How to find the sum of digits of a positive number using recursion?
func main() {
	fmt.Printf("Sum of digits in 125 is %d\n", getSumOfDigitsByRecursion(125))
	fmt.Printf("Sum of digits in 123 is %d\n", getSumOfDigitsByRecursion(123))
	fmt.Printf("Sum of digits in 1 is %d\n", getSumOfDigitsByRecursion(1))
	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("Sum of digits in 125 is %d\n", getSumOfDigitsByIteration(125))
	fmt.Printf("Sum of digits in 123 is %d\n", getSumOfDigitsByIteration(123))
	fmt.Printf("Sum of digits in 1 is %d\n", getSumOfDigitsByIteration(1))
}

// using recursion
func getSumOfDigitsByRecursion(num int) int {
	if num < 10 {
		return num
	}
	return num%10 + getSumOfDigitsByRecursion(num/10)
}

// using iteration
func getSumOfDigitsByIteration(num int) int {
	result := 0
	for num >= 10 {
		result += num % 10
		num = num / 10
	}
	result += num
	return result
}
