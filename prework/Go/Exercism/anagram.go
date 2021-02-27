package anagram

import "strings"

// Notes
//  Time complexity: O(n2)
//  A better alternative would have been create an array of 25 elements and increment each index based on letter
//  and then compare both arrays after converting them to string :)

func getASCIISum(s string) (sum int) {
	for _, c := range s {
		c = (c & 0xdf) - 'A'
		sum += int(c)
	}
	return
}

func getBitsetOfLetters(s string) (bs int32) {
	for _, c := range s {
		c = (c & 0xdf) - 'A'
		if c > 25 || c < 0 {
			continue
		}
		bs |= 1 << c
	}
	return
}

// Detect checks which of the provided candidates are anagrams of the subject. An anagram is a new word build from the rearrengement of the available letters
func Detect(subject string, candidates []string) (output []string) {
	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}
		if getASCIISum(subject) != getASCIISum(candidate) {
			continue
		}
		if getBitsetOfLetters(subject) != getBitsetOfLetters(candidate) {
			continue
		}
		if strings.ToLower(subject) == strings.ToLower(candidate) {
			continue
		}
		output = append(output, candidate)
	}
	return output
}
