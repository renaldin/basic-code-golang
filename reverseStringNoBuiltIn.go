package main

import (
	"fmt"
	"strings"
)

func reverseString(text string) string {
	textArray := strings.Split(text, "")
	var resultArray []string
	for _, t := range textArray {
		resultArray = append([]string{t}, resultArray...)
	}
	return strings.Join(resultArray, "")
}

// func main() {
// 	text := "hello11"
// 	resultReverse := reverseString(text)
// 	fmt.Println("Input: ", text)
// 	fmt.Println("Output: ", resultReverse)
}

// 2. Reverse String

// Buat fungsi yang membalikkan string tanpa menggunakan fungsi built-in seperti reverse().

// Contoh:

// Input: "hello"
// Output: "olleh"