package main

import "fmt"

func decimalToBiner(n int) {
	if n == 0 {
		return
	}

	decimalToBiner(n / 2)
	fmt.Print(n % 2)
}

// func main() {
// 	n := 4
// 	decimalToBiner(n)
// }

// case
// change decimal to biner