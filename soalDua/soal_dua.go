package soalDua

import (
	"fmt"
	"sort"
	"strings"
)

func Jawab(inputArr []string) {
	//membuat map
	str := strings.Join(inputArr, "")
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}

	//count  freqcount
	var keys []rune
	values := make(map[int]int)
	for key, value := range charMap {
		keys = append(keys, key)
		values[value]++
	}
	//sort char by frequency
	sort.SliceStable(keys, func(i, j int) bool {
		return charMap[keys[i]] > charMap[keys[j]]
	})

	//print by freq order then asci order
	var asciiSortStr []string
	flag := 1
	for _, char := range keys {
		flag = values[charMap[char]]
		if flag == 1 {
			if len(asciiSortStr) != 0 {
				asciiSortStr = append(asciiSortStr, string(char))
				sort.SliceStable(asciiSortStr, func(i, j int) bool {
					return asciiSortStr[i] < asciiSortStr[j]
				})

				newStr := strings.Join(asciiSortStr, "")

				fmt.Print(newStr)
				asciiSortStr = nil
			} else {
				fmt.Print(string(char))
			}
		} else {
			asciiSortStr = append(asciiSortStr, string(char))
			values[charMap[char]]--
		}
	}
	fmt.Println()
}
