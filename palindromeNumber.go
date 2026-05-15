package main

import (
	"strconv"
	"strings"
)

func palindromeNumber(num int) bool {
	str := strconv.Itoa(num)
	numArr := strings.Split(str, "")
	left := 0
	right := len(numArr) - 1
	for left < right {
		if numArr[left] != numArr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// func main() {
// 	num := 212212
// 	result := palindromeNumber(num)
// 	fmt.Println("Output:", result)
// }



// 1. Palindrome Number

// Diberikan sebuah integer x, tentukan apakah angka tersebut palindrome atau tidak.

// Palindrome adalah angka yang dibaca sama dari depan maupun belakang.

// Contoh:

// Input: 121  
// Output: true
// Input: 123  
// Output: false