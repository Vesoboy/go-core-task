package main

import (
	"fmt"
	"math/rand"
)

func RandomNumber(number chan int) {
	defer close(number)
	number <- rand.Intn(1000)
}

func main() {
	number := make(chan int)

	go RandomNumber(number)

	fmt.Println("Случайное число:", <-number)
}
