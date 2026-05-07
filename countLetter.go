package main

type CharCount struct {
	Char string `json:"char"`
	Count int `json:"count"`
}

// func main() {
// 	text := "aaabbcccaaaac"
// 	textArr := strings.Split(text, "")

// 	var charCount []CharCount
// 	tempChar := textArr[0]
// 	tempCount := 1
// 	for i := 1; i < len(textArr); i++ {
// 		if (tempChar == textArr[i]) {
// 			tempCount++
// 		} else {
// 			charCount = append(charCount, CharCount{Char: tempChar, Count: tempCount})
// 			tempCount = 1
// 			tempChar = textArr[i]
// 		}
// 	}
// 	charCount = append(charCount, CharCount{Char: tempChar, Count: tempCount})
// 	for _, char := range charCount {
// 		fmt.Println(char.Char, " = ", char.Count)
// 	}
// }
