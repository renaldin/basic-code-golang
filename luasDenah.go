package main

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var n int
// 	fmt.Fscan(in, &n)

// 	x := make([]int, n)
// 	y := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &x[i], &y[i])
// 	}

// 	area := 0
// 	for i := 0; i < n; i++ {
// 		j := (i + 1) % n
// 		area += x[i]*y[j] - x[j]*y[i]
// 	}

// 	if area < 0 {
// 		area = -area
// 	}

// 	fmt.Println(area / 2)
// }

// Soal Pemrograman #2
// Bayangkan sebuah denah rumah tampak atas. Anda diberikan koordinat kartesius pojok-pojok denah rumah tersebut, terurut mengitari rumah hingga kembali ke titik semula (searah jarum jam). Outputkan luas dari rumah tersebut. Diasumsikan sisi rumah selalu sejajar dengan Sumbu X atau Sumbu Y (tidak ada sisi rumah yang miring).

// Ilustrasi di bawah merupakan ilustrasi dari Contoh Input yang diberikan di bawah. Perlu diperhatikan, koordinat yang diberikan tidak selalu bermulai pada koordinat (0, 0). Pada input, baris pertama akan diinputkan banyak titik sudut bangun tersebut.

// soal pemrograman geometri
// Contoh Input

// 6
// 0 0
// 0 3
// 3 3
// 3 1
// 4 1
// 4 0
// Contoh Output

// 10