package anagram

import (
	"fmt"
)

func anagramGroup(words []string) [][]string {
	result := [][]string{}
	wordMap := map[int64][]string{}
	for _, word := range words {
		val := int64(1)
		for _, char := range word {
			val *= int64(char)
		}
		wordMap[val] = append(wordMap[val], word)
	}

	for _, val := range wordMap {
		result = append(result, val)
	}
	return result
}

func Run() {
	words := []string{
		"kita",
		"atik",
		"tika",
		"aku",
		"kia",
		"makan",
		"kua",
	}
	result := anagramGroup(words)
	fmt.Println(result)
}
