package main

import (
	"strings"
)

func parentheses(str string) bool {
	charMap := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}
	strArr := strings.Split(str, "")
	modStr := len(strArr) % 2
	if modStr != 0 {
		return false
	}
	var tempArr []string
	for _, val := range strArr {
		if val == "(" || val == "{" || val == "[" {
			tempArr = append(tempArr, val)
		} else {
			lenTmepArr := len(tempArr) - 1
			if charMap[val] != tempArr[lenTmepArr] {
				return false
			}
			tempArr = append(tempArr[:lenTmepArr], tempArr[lenTmepArr + 1:]...)
		}
	}
	return true
}

// func main() {
// 	str := "[({[()]})]"
// 	result := parentheses(str)
// 	fmt.Println(result)
// }

// 4. Valid Parentheses (Stack)

// Diberikan string yang hanya berisi karakter (, ), {, }, [, ], tentukan apakah string tersebut valid.

// Valid jika:

// Setiap kurung buka ditutup dengan jenis yang sama
// Urutan penutupan benar

// Contoh:

// Input: "({[]})"
// Output: true