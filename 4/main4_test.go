package main

import (
	"testing"
)

func TestFindSlice(t *testing.T) {
	slice1 := []string{"one", "two", "tree", "four", "five", "six"}
	slice2 := []string{"one", "two"}
	expected := []string{"tree", "four", "five", "six"}

	findSl := FindSlice(slice1, slice2)

	SliceMap := make(map[string]bool)
	for _, v := range findSl {
		SliceMap[v] = true
	}

	SliceMapExpected := make(map[string]bool)
	for _, v := range expected {
		SliceMapExpected[v] = true
	}

	for key, count1 := range SliceMap {
		if count2, exists := SliceMapExpected[key]; !exists || count1 != count2 {
			t.Errorf("Коллекция %s = %t, а не %t", key, count1, count2)
		}
	}

	if len(findSl) != len(expected) {
		t.Errorf("ожидалось %d элементов, а не %d", len(expected), len(findSl))
	}

}
