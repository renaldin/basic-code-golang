package main

import (
	"strings"
)

func frequencyLetter(text string) string {
	freq := map[string]int{}
	textLower := strings.ToLower(text)
	textArr := strings.Split(textLower, "")

	for _, t := range textArr {
		freq[t]++
	}

	maxCount := 0
	var textResult string

	for char, count := range freq {
		if count > maxCount {
			maxCount = count
			textResult = char
		}
	}

	return textResult
}

// func main() {
// 	text := "Hellooo"
// 	char := frequencyLetter(text)
// 	fmt.Println(char)
// }
