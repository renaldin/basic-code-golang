package main

import (
	"fmt"
	"strings"
)

func findFirstChar(text string) string {
	freq :=  map[string]int{}
	textSlice := strings.Split(text, "")
	for _, t := range textSlice {
		freq[t]++
	}
	for char, count := range freq {
		if count == 1 {
			return char
		} 
	}
	return "-1"
}

// func main() {
// 	text := "aabbcdd"
// 	result := findFirstChar(text)
// 	fmt.Println("Input :", text)
// 	fmt.Println("Output :", result)
// }


// 5. Find First Non-Repeating Character

// Diberikan sebuah string, temukan karakter pertama yang tidak berulang. Jika tidak ada, kembalikan null atau -1.

// Contoh:

// Input: "aabbcdd"  
// Output: "c"