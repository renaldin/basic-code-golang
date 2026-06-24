package main

import "fmt"

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		fizz := i % 3
		buzz := i % 5
		if fizz == 0 && buzz == 0 {
			fmt.Print("FizzBuzz ")
		} else if fizz == 0 {
			fmt.Print("Fizz ")
		} else if buzz == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

// func main() {
// 	n := 15
// 	fizzBuzz(n)
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