package main

import (
	"fmt"
)

func SumOfSquare(until int) int{
	sum := 0
	for i := 0; i <= until; i++ {
		sum += i*i
	}
	return sum
}

func SquareOfSum(until int) int{
	sum := 0
	for i := 0; i <= until; i++ {
		sum += i
	}
	sum = sum * sum
	return sum
}
func main() {
	fmt.Println(SquareOfSum(100) - SumOfSquare(100))
}