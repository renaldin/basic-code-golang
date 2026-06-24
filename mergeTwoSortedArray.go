package main

import (
	"sort"
)

func mergeArray(arr1 []int, arr2 []int) []int {
	for _, val := range arr2 {
		arr1 = append(arr1, val)
	}
	sort.Ints(arr1)
	return arr1
}

// func main() {
// 	arr1 := []int{1, 3, 5}
// 	arr2 := []int{2, 4, 6}
// 	result := mergeArray(arr1, arr2)
// 	fmt.Println(result)
// }


// 2. Merge Two Sorted Arrays

// Diberikan dua array yang sudah terurut ascending, gabungkan keduanya menjadi satu array yang tetap terurut.

// Contoh:

// Input:  
// arr1 = [1, 3, 5]  
// arr2 = [2, 4, 6]

// Output: [1, 2, 3, 4, 5, 6]