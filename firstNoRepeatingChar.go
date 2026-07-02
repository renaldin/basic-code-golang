package main

import (
	"strings"
)

func firstNoRepeatingChar(str string) string {
	freq := make(map[string]int)
	strArr := strings.Split(str, "")
	for _, value := range strArr {
		freq[value]++
	}
	var idx []int
	for key, val := range freq {
		if val == 1 {
			for index, value := range strArr {
				if value == key {
					idx = append(idx, index)
				}
			}
		}
	}
	if len(idx) == 0 {
		return ""
	}
	return strArr[idx[0]]
}

// func main() {
// 	str := "swwiiss"
// 	result := firstNoRepeatingChar(str)
// 	fmt.Println(result)
// }

// 6. Find the First Non-Repeating Character (Medium)

// Problem:
// Given a string, return the first character that appears only once. If no such character exists, return null.

// Example:

// Input:
// "swiss"

// Output:
// "w"