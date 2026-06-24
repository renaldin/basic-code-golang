package main

import "fmt"

func twoSum(nums []int, target int) {
	var arrCondition []int
	for key, val := range nums {
		arrCondition = append(arrCondition, key)
		for key2, val2 := range nums {
			found := false
			for _, valCondition := range arrCondition {
				if valCondition == key2 {
					found = true
				}
			}
			if found != true {
				counting := val + val2
				if counting == target {
					fmt.Println([]int{key, key2})
				}
			}
		}
	}
}

// func main() {
// 	nums := []int{2, 7, 11, 15, 4, 5, 7}
// 	target := 9
// 	twoSum(nums, target)
// }

// 1. Two Sum (Array & Hashing)

// Diberikan sebuah array integer nums dan sebuah integer target, kembalikan index dari dua angka yang jika dijumlahkan menghasilkan target.

// Ketentuan:

// Setiap input pasti memiliki tepat satu solusi
// Tidak boleh menggunakan elemen yang sama dua kali

// Contoh:

// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1]
