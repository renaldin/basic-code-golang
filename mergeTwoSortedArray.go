package main

import (
	"sort"
)

func mergeTwoArray(arr1 []int, arr2 []int) []int {
	arr := arr1
	for _, num := range arr2 {
		arr = append(arr, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

// func main() {
// 	arr1 := []int{1, 3, 5}
// 	arr2 := []int{2, 4, 6}
// 	result := mergeTwoArray(arr1, arr2)
// 	fmt.Println(result)
// }


// 2. Merge Two Sorted Arrays

// Diberikan dua array yang sudah terurut ascending, gabungkan keduanya menjadi satu array yang tetap terurut.

// Contoh:

// Input:  
// arr1 = [1, 3, 5]  
// arr2 = [2, 4, 6]

// Output: [1, 2, 3, 4, 5, 6]