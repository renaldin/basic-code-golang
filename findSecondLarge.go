package main

import (
	"sort"
)

func secondLarge(nums []int) int {
	sort.Ints(nums)
	secondLarga := len(nums) - 2
	return nums[secondLarga]
}

// func main() {
// 	nums := []int{10, 5, 8, 20, 15}
// 	result := secondLarge(nums)
// 	fmt.Println(result)
// }

// 3. Find the Second Largest Number

// Input:

// [10, 5, 8, 20, 15]

// Output:

// 15