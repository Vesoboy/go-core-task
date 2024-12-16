package main

import (
	"crypto/sha256"
	"fmt"
)

func TypeP[Tint, Tint8, Tint16, Tfloat64, Tstring, Tbool, Tcomplex64 any](
	x Tint, x8 Tint8, x16 Tint16, x64 Tfloat64, xstr Tstring, xbool Tbool, xcomplex Tcomplex64) {
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T\n", x, x8, x16, x64, xstr, xbool, xcomplex)
}

func StrFormat[Tint, Tint8, Tint16, Tfloat64, Tstring, Tbool, Tcomplex64 any](
	x Tint, x8 Tint8, x16 Tint16, x64 Tfloat64, xstr Tstring, xbool Tbool, xcomplex Tcomplex64) string {
	str := fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v\n", x, x8, x16, x64, xstr, xbool, xcomplex)

	fmt.Println(str)
	return str
}

func hash(str, salt string) string {
	runeText := []rune(str)
	data := []byte(string(runeText))
	var s = append(data, []byte(salt)...)

	hashResult := sha256.Sum256(s)
	hashString := fmt.Sprintf("%x", hashResult)

	fmt.Printf("Hash: %s", hashString)
	return hashString
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	TypeP(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	str := StrFormat(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	hash(str, "go-2024")
}
