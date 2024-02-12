package soalSatu

import (
	"fmt"
	"strings"
)

type CharCount struct {
	Char  string
	Count int
}

func Jawab(input string) {
	input = strings.ToLower(input)             // mengonversi ke huruf kecil
	input = strings.ReplaceAll(input, " ", "") // menghapus spasi " "

	var charCounts []CharCount
	for _, char := range input {
		charStr := string(char)
		found := false
		// Cek apakah karakter sudah ada dalam slice charCounts
		for i, cc := range charCounts {
			if cc.Char == charStr {
				charCounts[i].Count++
				found = true
				break
			}
		}
		// Jika karakter belum ada dalam slice charCounts, tambahkan ke slice
		if !found {
			charCounts = append(charCounts, CharCount{Char: charStr, Count: 1})
		}

	}

	// membuat format print
	var result []string
	for _, cc := range charCounts {
		result = append(result, fmt.Sprintf("%s=%d", cc.Char, cc.Count))
	}
	// print expect
	fmt.Println(strings.Join(result, ","))
}
