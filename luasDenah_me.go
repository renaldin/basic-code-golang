package main

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var n int
// 	fmt.Fscan(in, &n)

// 	x := make([]int, n) // [0, 4, 4, 0]
// 	y := make([]int, n) // [0, 0, 3, 3]

// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &x[i], &y[i])
// 	}

// 	area := 0
// 	for i := 0; i < n; i++ {
// 		j := (i + 1) % n
// 		area += (x[i] * y[j]) - (x[j] * y[i])
// 	}

// 	if area < 0 {
// 		area = -area
// 	}

// 	fmt.Println(area / 2)
// }