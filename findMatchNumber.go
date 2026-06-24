package main

func findNumber(nums []int) []int {
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}
	var result []int
	for key, val := range freq {
		if val > 1 {
			result = append(result, key)
		}
	}
	return result
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 2, 3, 1}
// 	result := findNumber(nums)
// 	fmt.Println(result)
// }

// 1. Find Duplicate Numbers (Easy)

// Problem Statement

// Given an array of integers, return all duplicate numbers in the array.

// Example

// Input:
// [1, 2, 3, 4, 2, 5, 1]

// Output:
// [1, 2]

// Requirements

// Return unique duplicate numbers only.
// The order of output does not matter.