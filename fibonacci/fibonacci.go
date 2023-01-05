package main

import "fmt"

func main() {
	fmt.Printf("Fibonacci series of 5: %v\n", getFibonacciSeries(5))
	fmt.Printf("Fibonacci series of 10: %v\n", getFibonacciSeries(10))
}

func getFibonacciSeries(num uint) []uint {
	series := []uint{0, 1}
	for i := 2; i < int(num); i++ {
		series = append(series, series[i-1]+series[i-2])
	}
	return series
}
