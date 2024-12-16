package main

import (
	"testing"
)

func TestTypeP(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	TypeP(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
}

func TestStrFormat(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	expected := "42, 42, 42, 3.14, Golang, true, (1+2i)\n"

	result := StrFormat(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	if result != expected {
		t.Errorf("StrFormat() = %v; want %v", result, expected)
	}
}

func TestHash(t *testing.T) {
	str := "Test string"
	salt := "go-2024"
	expectedHash := "1aae26ace3ca86cbb1233cd6eb7de352b78a45ecf715ac2e9d96f2a720194999"

	hashexpectedHash := hash(str, salt)

	if hashexpectedHash != expectedHash {
		t.Errorf("hash() = %v; want %v", hashexpectedHash, expectedHash)
	}
}
