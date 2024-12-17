package main

import (
	"fmt"
	"math"
	"math/rand"
)

func PowNumberTree(number chan float64, num float64) {
	var powNumber float64 = math.Pow(num, 3)
	number <- powNumber
	close(number)
}

func RandomNumbers() float64 {
	number := rand.Intn(10)
	fmt.Printf("Rand Value for Pow: %d\n", number)
	return float64(number)
}

func Merge(channels ...chan float64) <-chan float64 {
	out := make(chan float64)
	sem := make(chan struct{}, len(channels))

	for _, value := range channels {
		sem <- struct{}{}
		go func(c chan float64) {
			defer func() { <-sem }()
			for val := range c {
				out <- val
			}

		}(value)
	}

	go func() {
		for range channels {
			sem <- struct{}{}
		}
		close(out)
	}()

	return out

}

func main() {

	number1 := make(chan float64)
	number2 := make(chan float64)
	number3 := make(chan float64)

	go PowNumberTree(number1, RandomNumbers())
	go PowNumberTree(number2, RandomNumbers())
	go PowNumberTree(number3, RandomNumbers())

	merged := Merge(number1, number2, number3)

	for val := range merged {
		fmt.Printf("Rand Value: %.2f\n", val)
	}
}
