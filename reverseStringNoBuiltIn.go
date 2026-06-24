package main

import (
	"strings"
)

func reverseString(str string) string {
	strArr := strings.Split(str, "")
	right := len(strArr) - 1
	var resultTemp []string
	for i := 0; i < len(strArr); i++ {
		resultTemp = append(resultTemp, strArr[right])
		right--
	}
	return strings.Join(resultTemp, "")
}

// func main() {
// 	str := "helloo"
// 	result := reverseString(str)
// 	fmt.Println(result)
// }

// 2. Reverse String

// Buat fungsi yang membalikkan string tanpa menggunakan fungsi built-in seperti reverse().

// Contoh:

// Input: "hello"
// Output: "olleh"