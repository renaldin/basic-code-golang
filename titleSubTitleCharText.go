package main

import (
	"strings"
	"unicode"
)

func toProperCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func convertText(text string) (string, string) {
	words := strings.Fields(text)
	words[0] = toProperCase(words[0])
	for i := 1; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return toProperCase(text), strings.Join(words, " ")
}

// func main() {
// 	text := "SeLamAT PAGi semua halOo"
// 	text1, text2 := convertText(text)
// 	fmt.Println("Title format :", text1)
// 	fmt.Println("Subtitle format :", text2)
// }


// Soal Pemrograman #1
// Diberikan sebuah string yang dapat mengandung huruf (kapital ataupun non-kapital) dan karakter spasi di dalamnya, ubah format string tersebut menjadi format penulisan judul dan format penulisan biasa.

// Format penulisan judul menjadikan huruf pertama setiap katanya menjadi kapital sedangkan huruf lainnya tidak, sedangkan format penulisan biasa hanya menjadikan huruf pertama string menjadi kapital dan huruf lainnya tidak.

// Contoh Input:

// SeLamAT PAGi semua halOo
// Contoh Output:

// Format Judul: Selamat Pagi Semua Haloo
// Format Biasa: Selamat pagi semua haloo