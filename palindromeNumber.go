package main

import (
	"strconv"
	"strings"
)

func isPalindromeNumber(number int) bool {
	str := strconv.Itoa(number)
	arrStr := strings.Split(str, "")
	left := 0
	right := len(arrStr) - 1
	for left < right {
		if arrStr[left] != arrStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// func main() {
// 	number := 12121
// 	result := isPalindromeNumber(number)
// 	fmt.Println(result)
// }



// 1. Palindrome Number

// Diberikan sebuah integer x, tentukan apakah angka tersebut palindrome atau tidak.

// Palindrome adalah angka yang dibaca sama dari depan maupun belakang.

// Contoh:

// Input: 121  
// Output: true
// Input: 123  
// Output: false