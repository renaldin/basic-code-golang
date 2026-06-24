package main

import (
	"fmt"
	"strings"
)

func isPalindromeChekcer(str string) bool {
	strNew := strings.ReplaceAll(str, " ", "")
	strLower := strings.ToLower(strNew)
	strArr := strings.Split(strLower, "")
	left := 0 
	right := len(strArr) - 1
	fmt.Println(left, right)
	for left < right {
		fmt.Println(strArr[left], strArr[right])
		if strArr[left] != strArr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := "A man a plan a canal Panama"
	result := isPalindromeChekcer(str)
	fmt.Println(result)
}

// Problem:
// Write a function that determines whether a given string is a palindrome. Ignore spaces, punctuation, and letter casing.

// Example:

// Input:
// "A man a plan a canal Panama"

// Output:
// true