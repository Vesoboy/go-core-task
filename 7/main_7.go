package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
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
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c chan float64) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	go func() {
		wg.Wait()
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
	sumMer := 0.0
	for val := range merged {
		sumMer += val
		fmt.Printf("Rand Value: %.2f\n", val)
	}

	fmt.Printf("Sum: %.2f\n", sumMer)

}
