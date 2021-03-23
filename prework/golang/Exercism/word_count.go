package wordcount

import (
	"strings"
	"unicode"
)

// Notes:
//  - It performs better to call toLower than the map

// sentence = strings.Map(func(r rune) rune {
// 	if unicode.IsUpper(r) {
// 		return r | 0x20
// 	}
// 	return r
// }, sentence)

// Frequency type handles result
type Frequency map[string]int

// WordCount checks the number of times a word is repeated in provided sentence
func WordCount(sentence string) Frequency {
	countMap := make(Frequency)
	sentence = strings.ToLower(sentence)
	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	for _, word := range words {
		word = strings.Trim(word, "'")
		countMap[word]++
	}

	return countMap
}
