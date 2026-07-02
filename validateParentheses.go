package main

import (
	"strings"
)

func validateParentheses(str string) bool {
	mapChar := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	strArr := strings.Split(str, "")
	var tempStr string
	for _, val := range strArr {
		if val == "(" || val == "[" || val == "{" {
			tempStr = val
		} else {
			if tempStr != mapChar[val] {
				return false
			}
		}
	}
	return true
}

// func main() {
// 	str := "()[]{}"
// 	result := validateParentheses(str)
// 	fmt.Println(result)
// }

// 8. Validate Parentheses (Medium)

// Problem:
// Given a string containing only the characters (, ), {, }, [, and ], determine whether the input string is valid.

// A string is valid if:

// Every opening bracket has a matching closing bracket.
// Brackets are closed in the correct order.

// Example 1

// Input:
// "()[]{}"

// Output:
// true

// Example 2

// Input:
// "(]"

// Output:
// false