package main

import (
	"strings"
)

func countChar(str string) map[string]int {
	freq := make(map[string]int)
	arrStr := strings.Split(str, "")
	for _, val := range arrStr {
		freq[val]++
	}
	return freq
}

// func main() {
// 	str := "programmer"
// 	result := countChar(str)
// 	fmt.Println(result)
// }

// 2. Count Character Occurrences

// Input:

// "programmer"

// Output:

// {
//   "p": 1,
//   "r": 3,
//   "o": 1,
//   "g": 1,
//   "a": 1,
//   "m": 2,
//   "e": 1
// }