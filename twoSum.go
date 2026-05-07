package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, found := seen[complement]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

// func main() {
// 	nums := []int{2, 1, 7, 15}
// 	target := 9
// 	arrResult := twoSum(nums, target)
// 	fmt.Println(arrResult)
// }

// 1. Two Sum (Array & Hashing)

// Diberikan sebuah array integer nums dan sebuah integer target, kembalikan index dari dua angka yang jika dijumlahkan menghasilkan target.

// Ketentuan:

// Setiap input pasti memiliki tepat satu solusi
// Tidak boleh menggunakan elemen yang sama dua kali

// Contoh:

// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1]
