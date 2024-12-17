package main

import (
	"testing"
)

func TestPowNumberTree(t *testing.T) {
	number1 := make(chan []int)

	go func() {
		number1 <- []int{1, 2, 3, 4}
		close(number1)
	}()
	number2 := make(chan float64)
	expected := []float64{1, 8, 27, 64}

	go PowNumberTree(number1, number2)

	for i := 0; i < 4; i++ {
		if val := <-number2; val != expected[i] {
			t.Errorf("PowNumberTree() = %v; а должено быть %v", val, expected[i])
		}
	}
}

func TestGenerateNumbers(t *testing.T) {
	number1 := make(chan []int)

	go func() {
		for i := 0; i < 10000; i++ {
			GenerateNumbers(number1)
		}
		close(number1)
	}()

	expected := 255.0 // uint8 от 0 до 255

	for val := range number1 {
		for _, num := range val {
			if val := float64(num); val > expected {
				t.Errorf("GenerateNumbers() = %v; uint8 от 0 до %v", val, expected)
			}
		}
	}
}
