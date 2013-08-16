package main

import (
    "fmt"
)

func SumOfMultiples(factor int) int {
    sum := 0
    for i := 0; i < 10000; i++ {
        if i % factor == 0 {
            sum += i
        }
    }
    return sum
}

func main() {
    fmt.Println(SumOfMultiples(5) + SumOfMultiples(3) - SumOfMultiples(15))
}