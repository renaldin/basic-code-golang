package main

import (
	"strings"
)

func coutningLetter(str string) string {
	freq := make(map[string]int)
	strToLower := strings.ToLower(str)
	strArr := strings.Split(strToLower, "")
	for _, str := range strArr {
		freq[str]++
	}
	var result string
	maxCount := 0
	for key, val := range freq {
		if val > maxCount {
			maxCount = val
			result = key
		}
	}
	return result
}

// func main() {
// 	str := "Hellooo"
// 	result := coutningLetter(str)
// 	fmt.Println(result)
// }

// buatkan function yang dimana outoutnya adalah jumlah huruf yang paling besar yang dikeluarkan
// contoh input hellooo dan outputnya adalah o dan harus mengeluarkan huruf kecil