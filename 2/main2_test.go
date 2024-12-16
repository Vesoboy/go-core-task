package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(slice)
	if len(result) != len(expected) {
		t.Errorf("sliceExample() = %v; want %v", result, expected)
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := copySlice(slice)
	slice = append(slice, 323) //проверка на то, что оригинальный слайс не влияет на копию
	if len(result) != len(expected) {
		t.Errorf("copySlice() = %v; want %v", result, expected)
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := addElements(slice, 11)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("addElements() = %v; want %v", result, expected)
		}
	}
	if len(result) != len(expected) {
		t.Errorf("addElements() = %v; want %v", result, expected)
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	result := removeElement(slice, 5)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("removeElement() = %v; want %v", result, expected)
		}
	}
	if len(result) != len(expected) {
		t.Errorf("removeElement() = %v; want %v", result, expected)
	}
}
