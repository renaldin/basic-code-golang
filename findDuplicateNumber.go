package main

func findDuplicate(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
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
// 	nums := []int{1, 2, 3, 4, 2, 5, 3}
// 	result := findDuplicate(nums)
// 	fmt.Println(result)
// }

// 3. Find Duplicate Number

// Diberikan array integer, cari angka yang muncul lebih dari satu kali.

// Contoh:

// Input: [1, 2, 3, 4, 2, 5]  
// Output: 2