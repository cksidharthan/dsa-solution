package main

import "fmt"

func main() {
	fmt.Printf("Factorial of 5: %d\n", getFactorial(5))
	fmt.Printf("Factorial of 10: %d\n", getFactorial(10))
	fmt.Printf("Factorial of 0: %d\n", getFactorial(0))
}

func getFactorial(number uint) uint {
	if number < 2 {
		return 1
	}
	return getFactorial(number-1) * number
}
