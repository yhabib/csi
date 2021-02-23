package isogram

import (
	"unicode"
)

// Notes:
// Sets are creted in Go as mab of rune to bool: make(map[rune]bool)
// Performancewise is better to operate at the rune level, faster
// More performant to iterate through the list than the Mao
//  eg: if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {

// IsIsogram checks if a word is an isogram: no repeated letters
func IsIsogram(s string) bool {
	seenKeys := make(map[rune]bool)

	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		lowC := unicode.ToLower(c)
		if seenKeys[lowC] {
			return false
		}
		seenKeys[lowC] = true
	}
	return true
}
