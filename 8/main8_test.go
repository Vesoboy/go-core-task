package main

import (
	"testing"
)

func TestPowNumberTree(t *testing.T) {
	number := make(chan float64)
	expected := 8.0
	go PowNumberTree(number, 2)

	if val := <-number; val != expected {
		t.Errorf("PowNumberTree() = %v; want %v", val, expected)
	}
}

func TestMerge(t *testing.T) {
	number1 := make(chan float64)
	number2 := make(chan float64)
	number3 := make(chan float64)

	go PowNumberTree(number1, 2)
	go PowNumberTree(number2, 3)
	go PowNumberTree(number3, 4)

	merged := Merge(number1, number2, number3)
	expected := 99.0

	sumMer := 0.0
	for val := range merged {
		sumMer += val
	}

	if val := sumMer; val != expected {
		t.Errorf("Сумма Merge() = %v; в не %v", val, expected)
	}
}
