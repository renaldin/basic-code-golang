package main

import (
	"fmt"
	"strings"
)

func extractString(str string) {
	strArr := strings.Split(str, " ")
	var resultArrTitle []string
	var resultArrSubTitle []string
	for keyFirst, val := range strArr {
		valArr := strings.Split(val, "")
		for key, valTwo := range valArr {
			if key == 0 {
				resultArrTitle = append(resultArrTitle, strings.ToUpper(valTwo))
			} else {
				resultArrTitle = append(resultArrTitle, strings.ToLower(valTwo))
			}
			if keyFirst == 0 && key == 0 {
				resultArrSubTitle = append(resultArrSubTitle, strings.ToUpper(valTwo))
			} else {
				resultArrSubTitle = append(resultArrSubTitle, strings.ToLower(valTwo))
			}
		}
		resultArrTitle = append(resultArrTitle, " ")
		resultArrSubTitle = append(resultArrSubTitle, " ")
	}
	fmt.Println("Title Format:", strings.Join(resultArrTitle, ""))
	fmt.Println("Subtitle Format:", strings.Join(resultArrSubTitle, ""))
}

// func main() {
// 	str := "SeLamAT PAGi semua halOo"
// 	extractString(str)
// }


// Soal Pemrograman #1
// Diberikan sebuah string yang dapat mengandung huruf (kapital ataupun non-kapital) dan karakter spasi di dalamnya, ubah format string tersebut menjadi format penulisan judul dan format penulisan biasa.

// Format penulisan judul menjadikan huruf pertama setiap katanya menjadi kapital sedangkan huruf lainnya tidak, sedangkan format penulisan biasa hanya menjadikan huruf pertama string menjadi kapital dan huruf lainnya tidak.

// Contoh Input:

// SeLamAT PAGi semua halOo
// Contoh Output:

// Format Judul: Selamat Pagi Semua Haloo
// Format Biasa: Selamat pagi semua haloo