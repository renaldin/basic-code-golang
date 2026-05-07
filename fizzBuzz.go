package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		kelip3 := i % 3
		kelip5 := i % 5
		if kelip3 == 0 && kelip5 == 0 {
			result = append(result, "FizzBuzz")
		} else if kelip3 == 0 {
			result = append(result, "Fizz")
		} else if kelip5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result

}

// func main() {
// 	n := 15
// 	result := fizzBuzz(n)
// 	fmt.Println("Input: n = ", n)
// 	fmt.Println("Output: ", result)
// }

// 3. FizzBuzz

// Tulis program yang mencetak angka dari 1 sampai n, dengan aturan:

// Kelipatan 3 → cetak "Fizz"
// Kelipatan 5 → cetak "Buzz"
// Kelipatan 3 dan 5 → cetak "FizzBuzz"
// Selain itu → cetak angka itu sendiri

// Contoh:

// Input: n = 5
// Output: [1, 2, "Fizz", 4, "Buzz"]