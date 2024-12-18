package main

import (
	"testing"
)

func TestPowNumberTree(t *testing.T) {
	number := make(chan []int)

	go func() {
		number <- []int{1, 2, 3, 4}
		close(number)
	}()
	numberPow := make(chan float64)
	expected := []float64{1, 8, 27, 64}

	go PowNumberTree(number, numberPow)

	for i := 0; i < 4; i++ {
		if val := <-numberPow; val != expected[i] {
			t.Errorf("PowNumberTree() = %v; а должено быть %v", val, expected[i])
		}
	}
}

func TestGenerateNumbers(t *testing.T) {
	number := make(chan []int)
	countNumber := 1000

	go GenerateNumbers(number, countNumber)

	expected := 255.0 // uint8 от 0 до 255

	for val := range number {
		for _, num := range val {
			if val := float64(num); val > expected {
				t.Errorf("GenerateNumbers() = %v; uint8 от 0 до %v", val, expected)
			}
		}
	}
}
