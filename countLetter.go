package main

type CharCount struct {
	Char string `json:"char"`
	Count int `json:"count"`
}

// func main() {
// 	text := "aaabbcccaaaac"
// 	textArr := strings.Split(text, "")

// 	tempChar := textArr[0]
// 	tempCount := 1
// 	charCount := []CharCount{}

// 	for i := 1; i < len(textArr); i++{
// 		if tempChar == textArr[i] {
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


// input
// "aaabbcccaaaac"

// output
// a  =  3
// b  =  2
// c  =  3
// a  =  4
// c  =  1