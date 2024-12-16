package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() *StringIntMap {
	newMap := NewStringIntMap()
	for key, value := range m.data {
		newMap.data[key] = value
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	if _, ok := m.data[key]; ok {
		return true
	} else {
		return false
	}
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}

func (m *StringIntMap) GetAll() {
	for key, value := range m.data {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func main() {
	newMap := NewStringIntMap()
	fmt.Print("Выведем все что записали: \n")
	newMap.Add("one", 41)
	newMap.Add("two", 42)
	newMap.Add("tree", 43)
	newMap.Add("four", 44)
	newMap.Add("five", 45)
	newMap.Add("six", 46)
	newMap.GetAll()

	fmt.Print("\nУдалим элемент six: ")
	fmt.Print("\nВыведем что получилось: \n")
	newMap.Remove("six")
	newMap.GetAll()

	newMapCopy := newMap.Copy()
	fmt.Printf("\nВыведем что получилось после копирования: \n")
	newMapCopy.GetAll()

	boolKeyExists := newMap.Exists("one")
	fmt.Printf("\nСделаем проверку на существование one")
	fmt.Printf("\none существует: %t\n", boolKeyExists)

	fmt.Printf("\nПолучим значение tree: ")
	if value, ok := newMap.Get("tree"); ok {
		fmt.Printf("\nKey: %s, Value: %d\n", "tree", value)
	}
}
