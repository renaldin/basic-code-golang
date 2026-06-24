package main

import (
	"strings"
)

func findFirstChar(str string) string {
	strArr := strings.Split(str, "")
	freq := make(map[string]int)

	for _, val := range strArr {
		freq[val]++
	}

	var resultArr []string
	for key, val := range freq {
		if val == 1 {
			resultArr = append(resultArr, key)
		}
	}

	if len(resultArr) == 0 {
		return "-1"
	}

	return resultArr[0]
}

// func main() {
// 	str := "aabbcd"
// 	result := findFirstChar(str)
// 	fmt.Println(result)
// }


// 5. Find First Non-Repeating Character

// Diberikan sebuah string, temukan karakter pertama yang tidak berulang. Jika tidak ada, kembalikan null atau -1.

// Contoh:

// Input: "aabbcdd"  
// Output: "c"