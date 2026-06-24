package main

import (
	"strings"
)

func romanToInteger(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sArry := strings.Split(s, "")
	lenS := len(s)
	result := 0
	for i := 0; i < lenS; i++ {
		current := romanMap[sArry[i]]
		if i + 1 < lenS {
			next := romanMap[sArry[i + 1]]
			if current < next {
				result -= current
			} else {
				result += current
			}
		} else {
			result += current
		}
	}
	return result
}

// func main() {
// 	roman := "MCMXCIV"
// 	result := romanToInteger(roman)
// 	fmt.Println(result)
// }


// 5. Roman to Integer

// Konversikan angka romawi menjadi integer.

// Contoh:

// Input: "IX"  
// Output: 9
// Input: "MCMXCIV"  
// Output: 1994