package main

import (
	"strings"
)

func isPalindrom(text string) string {
	arrText := strings.Split(text, "")

	left := 0
	right := len(arrText) - 1

	for left < right {
		if arrText[left] != arrText[right] {
			return "Not Palindrom"
		}
		left++
		right--
	}
	return "Palindrom"
}

// func main() {
// 	text := "level level levell"
// 	result := isPalindrom(text)
// 	fmt.Println(result)
// }