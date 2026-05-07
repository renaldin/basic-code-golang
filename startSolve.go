package main

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	fmt.Print("Input: ")
// 	var n int
// 	fmt.Fscan(in, &n)

// 	fmt.Println("Output:")
// 	for i := 1; i <= n; i++ {
// 		num1 := i + 1
// 		num2 := num1 + 1
// 		for j := 1; j <= n+3; j++ {
// 			if j == num1 {
// 				fmt.Print("*")
// 			} else if j == num2 {
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print(j)
// 			}
// 		}
// 		fmt.Println()
// 	}
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