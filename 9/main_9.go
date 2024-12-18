package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func PowNumberTree(number chan []int, numberPow chan<- float64) {
	defer close(numberPow)
	for val := range number {
		for _, num := range val {
			numberPow <- math.Pow(float64(num), 3)
		}
	}
}

func GenerateNumbers(number chan []int, count int) {
	defer close(number)
	for i := 0; i < count; i++ {
		number <- []int{rand.IntN(256)}
	}
}

func main() {

	number := make(chan []int)
	numberPow := make(chan float64)
	countNumber := 10
	go PowNumberTree(number, numberPow)

	go GenerateNumbers(number, countNumber)

	for val := range numberPow {
		fmt.Printf("Rand Value: %.2f\n", val)
	}
}
