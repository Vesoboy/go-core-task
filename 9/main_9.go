package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func PowNumberTree(number chan []int, number2 chan<- float64) {
	go func() {

		for val := range number {
			for _, num := range val {
				number2 <- math.Pow(float64(num), 3)
			}
		}
		close(number2)
	}()
}

func GenerateNumbers(ch chan []int) {
	numbers := []int{rand.IntN(256)}
	ch <- numbers
}

func main() {

	number1 := make(chan []int)
	number2 := make(chan float64)
	go PowNumberTree(number1, number2)

	go func() {
		for i := 0; i < 10; i++ {
			GenerateNumbers(number1)
		}
		close(number1)
	}()

	for val := range number2 {
		fmt.Printf("Rand Value: %.2f\n", val)
	}

}
