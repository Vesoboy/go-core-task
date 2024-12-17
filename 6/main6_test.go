package main

import (
	"testing"
)

func TestRandomNumber(t *testing.T) {
	number := make(chan int)

	go RandomNumber(number)

	randomValue, ok := <-number
	if !ok {
		t.Fatalf("Не успели получить случайное число")
	}

	// Проверяем, что число находится в ожидаемом диапазоне
	if randomValue < 0 || randomValue >= 1000 {
		t.Errorf("Значение не в диапазоне [0, 999], получено: %d", randomValue)
	}
}
