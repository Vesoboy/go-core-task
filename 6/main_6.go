package main

import (
	"fmt"
	"math/rand"
)

func RandomNumber(number chan int, count int) chan int {
	defer close(number)
	for i := 0; i < count; i++ {
		number <- rand.Intn(1000)
	}
	return number
}

func main() {

	number := make(chan int)
	countNumber := 10

	go RandomNumber(number, countNumber)

	for val := range number {
		fmt.Printf("Случайное число: %d\n", val)
	}
}
