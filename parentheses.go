package main

import (
	"fmt"
	"strings"
)

func parenttheses(text string) bool {
	var textArrayTwo []string
	mapDef := map[string]string {
		"}": "{",
		")": "(",
		"]": "[",
	}
	textArray := strings.Split(text, "")
	checkMod := len(textArray) % 2
	if checkMod != 0 {
		return false
	}
	for _, t := range textArray {
		if t == "(" || t == "{" || t == "[" {
			textArrayTwo = append(textArrayTwo, t)
		} else {
			lastIndex := len(textArrayTwo) - 1
			if lastIndex < 0 {
				return false
			}
			lastValue := textArrayTwo[lastIndex]
			if lastValue != mapDef[t] {
				return false
			}
			textArrayTwo = append(textArrayTwo[:lastIndex], textArrayTwo[lastIndex + 1:]...)
		}
	}
	return true
}

func main() {
	text := "{[]}"
	result := parenttheses(text)
	fmt.Println("Input : ", text)
	fmt.Println("Output : ", result)
}

// 4. Valid Parentheses (Stack)

// Diberikan string yang hanya berisi karakter (, ), {, }, [, ], tentukan apakah string tersebut valid.

// Valid jika:

// Setiap kurung buka ditutup dengan jenis yang sama
// Urutan penutupan benar

// Contoh:

// Input: "({[]})"
// Output: true