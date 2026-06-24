package main

import "fmt"

func startSolve(n int) {
	nj := n + 3
	startOne := 2
	startTwo := 3
	for i := 1; i <= n; i++ {
		for j := 1; j <= nj; j++ {
			if j == startOne || j == startTwo {
				fmt.Print("*")
			} else {
				fmt.Print(j)
			}
		}
		startOne++
		startTwo++
		fmt.Println()
	}
}

// func main() {
// 	n := 4
// 	startSolve(n)
// }

// Buatlah fungsi yang dapat menghasilkan output seperti dibawah ini:

// input n = 5
// 1**45678
// 12**5678
// 123**678
// 1234**78
// 12345**8

// input n = 4
// 1**4567
// 12**567
// 123**67
// 1234**7