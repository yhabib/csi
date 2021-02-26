package pangram

import (
	"unicode"
)

// Notes:
//  O(n) = 2*n = n, fixed to 26

// IsPangram checks if provided string contains all letter of the alphabet at least once
func IsPangram(input string) bool {
	var containsLetter [26]bool
	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}
		normalizedLowerCaseRune := unicode.ToLower(c) - 'a'
		containsLetter[normalizedLowerCaseRune] = true
	}
	for _, v := range containsLetter {
		if !v {
			return false
		}
	}

	return true
}

// Notes:
//  32bit integer as a bit set -> pangram if bs has first four bits to 0 and then all to 1
//  s[i] & 0xdf -> lower case into upper case, no change to upper case. ASCII distance from a -> A is a bit in the middle: 0xdf == 11011111
//  - 'A' to normalize it to 0
//  bs |= 1 << c first shifts 1 c positions to the left, eg: h -> 7 -> 10000000 and then does the bs |= to put that bit to one
//
// func IsPangram(s string) bool {
// 	var bs int32
// 	for i := 0; i < len(s); i++ {
// 		c := (s[i] & 0xdf) - 'A'
// 		if c > 25 || c < 0 {
// 			continue
// 		}
// 		bs |= 1 << c
// 	}
// 	return bs == 0x3ffffff
// }
