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
// 	n := 3
// 	decimalToBiner(n)
// }
