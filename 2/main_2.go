package main

import (
	"fmt"
	"math/rand"
)

func RandSlice(size int) []int {
	var randSlice []int
	for i := 0; i < size; i++ {
		randSlice = append(randSlice, rand.Intn(100))
	}
	return randSlice
}

func PrintSlice(slice []int) {
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
}

func sliceExample(slice []int) []int {
	var newSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func addElements(slice []int, n int) []int {
	slice = append(slice, n)
	return slice
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	newSlice := append(slice[:index], slice[index+1:]...)
	return newSlice
}

func main() {
	//originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	originalSlice := RandSlice(10)
	fmt.Print("Вывод слайса, который мы сгенерировали: ")
	PrintSlice(originalSlice)

	newSlice := sliceExample(originalSlice)
	fmt.Print("\nВывод слайса, который содержит только четные числа: ")
	PrintSlice(newSlice)

	newSlice = addElements(newSlice, rand.Intn(100))
	fmt.Print("\nВывод слайса, после добавления нового элемента: ")
	PrintSlice(newSlice)

	newSlice = copySlice(originalSlice)
	fmt.Print("\nВывод слайса, после копирования: ")
	PrintSlice(newSlice)

	randIndexSlice := rand.Intn(len(newSlice))
	fmt.Printf("\n\nУдалим из слайса элемент с индесом: \t%d ", randIndexSlice)
	fmt.Print("\nВывод слайса, до удаления элемента: \t")
	PrintSlice(newSlice)
	fmt.Print("\nВывод слайса, после удаления элемента: \t")
	newSlice = removeElement(newSlice, randIndexSlice)
	PrintSlice(newSlice)
}
