package main

import (
	"sort"
)

func sortedTwoArray(arr1 []int, arr2 []int) []int {
	for _, val := range arr2 {
		arr1 = append(arr1, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr1)))
	return arr1
}

// func main() {
// 	arr1 := []int{1, 3, 5}
// 	arr2 := []int{2, 4, 6}
// 	result := sortedTwoArray(arr1, arr2)
// 	fmt.Println(result)
// }

// 7. Merge Two Sorted Arrays (Medium)

// Problem:
// Write a function that merges two sorted arrays into one sorted array.

// Example:

// Input:
// arr1 = [1, 3, 5]
// arr2 = [2, 4, 6]

// Output:
// [1, 2, 3, 4, 5, 6]