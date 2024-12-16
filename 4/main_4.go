package main

import (
	"fmt"
)

func FindSlice(slice1, slice2 []string) []string {
	var result []string

	newSlice2Map := make(map[string]bool)
	for _, v := range slice2 {
		newSlice2Map[v] = true
	}

	for _, v := range slice1 {
		if _, ok := newSlice2Map[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"fig", "21", "date", "lead", "gno1"}
	clF := FindSlice(slice1, slice2)
	fmt.Printf("%v ", clF)

}
