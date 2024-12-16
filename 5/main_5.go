package main

import (
	"fmt"
)

func FindSliceInt(slice1, slice2 []int) (bool, []int) {
	var result []int
	var findBool bool = false
	newSlice2Map := make(map[int]bool)
	for _, v := range slice2 {
		newSlice2Map[v] = true
	}

	for _, v := range slice1 {
		if _, ok := newSlice2Map[v]; ok {
			result = append(result, v)
			findBool = true
		}
	}
	return findBool, result
}

func main() {
	a := []int{653, 3, 3, 58, 678, 64}
	b := []int{642, 2, 3, 43, 3, 1}

	if ok, newSl := FindSliceInt(a, b); ok {
		fmt.Println(ok, newSl)
	} else {
		fmt.Println(ok, newSl)
	}

}
