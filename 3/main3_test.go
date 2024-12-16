package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)
	m.Add("key2", 20)

	if len(m.data) != 2 {
		t.Errorf("Ожидалось 2 элемента, а не %d", len(m.data))
	}

	if val, ok := m.data["key1"]; !ok || val != 10 {
		t.Errorf("ожидался key1 = 10, а не %d", val)
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)
	m.Remove("key1")

	if _, ok := m.data["key1"]; ok {
		t.Errorf("ожидался что key1 удалится, но увы")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)
	m.Add("key2", 20)

	copyMap := m.Copy()
	if len(copyMap.data) != 2 {
		t.Errorf("в скопираванном map ожидалось 2 элемента, а не %d", len(copyMap.data))
	}

	if val, ok := copyMap.data["key1"]; !ok || val != 10 {
		t.Errorf("в скопираванном map ожидался key1 = 10, а не %d", val)
	}

	m.Remove("key1")
	if _, ok := copyMap.data["key1"]; !ok {
		t.Errorf("ожидался что key1 удалится в оригинальном map, но увы")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)

	if !m.Exists("key1") {
		t.Errorf("key1 существует, но видимо в другой реальности")
	}

	if m.Exists("key2") {
		t.Errorf("key2 не существует")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)

	val, ok := m.Get("key1")
	if !ok || val != 10 {
		t.Errorf("key1 = 10, а не %d", val)
	}

	_, ok = m.Get("key2")
	if ok {
		t.Errorf("key2 не существует")
	}
}
