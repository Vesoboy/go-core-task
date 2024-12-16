package main

import (
	"testing"
)

func TestFindSliceInt(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 32, 123, 1}
	expected := []int{1, 2, 3}
	expectedBool := true

	ok, findSl := FindSliceInt(a, b)

	SliceMap := make(map[int]bool)
	for _, v := range findSl {
		SliceMap[v] = true
	}

	SliceMapExpected := make(map[int]bool)
	for _, v := range expected {
		SliceMapExpected[v] = true
	}

	for key, count1 := range SliceMap {
		if count2, exists := SliceMapExpected[key]; !exists || count1 != count2 {
			t.Errorf("Коллекция %d = %t, а не %t", key, count1, count2)
		}
	}

	if len(findSl) != len(expected) {
		t.Errorf("ожидалось %d элементов, а не %d", len(expected), len(findSl))
	}

	if ok != expectedBool {
		t.Errorf("ожидалось %t, а не %t", expectedBool, ok)
	}

}
