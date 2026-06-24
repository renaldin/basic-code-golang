package main

import (
	"strings"
)

func isPalindrome(str string) string {
	arrStr := strings.Split(str, "")
	left := 0
	right := len(arrStr) - 1
	for left < right {
		if arrStr[left] != arrStr[right] {
			return "Not Palindrome"
		}
		left++
		right--
	}
	return "Palindrome"
}

// func main() {
// 	str := "level level"
// 	result := isPalindrome(str)
// 	fmt.Println(result)
// }