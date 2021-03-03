package lsproduct

import "errors"

// NOTES:
//  - Brute force: 0 -> len(digits), i -> span => O(n2) 1899699               638 ns/op              64 B/op          4 allocs/op
//  - Life multiplication: => O(n)											1057630              1144 ns/op              64 B/op          4 allocs/op
// 	Because of the size of N it performs better the quadratic solution

func hasNonDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return true
		}
	}
	return false
}

// LargestSeriesProduct calculates the largest product in the provided serie
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return -1, errors.New("span has to be smaller than length")
	}
	if span < 0 {
		return -1, errors.New("span has to be a positive integer")
	}
	if span == 0 {
		return 1, nil
	}
	if hasNonDigit(digits) {
		return -1, errors.New("digits can only contain numbers")
	}
	// Brute force
	// for i := 0; i <= len(digits)-span; i++ {
	// 	tempMax := 1
	// 	for j := 0; j < span; j++ {
	// 		tempMax *= int(digits[i+j] - '0')
	// 	}
	// 	if tempMax > max {
	// 		max = tempMax
	// 	}
	// }

	product, lenSb, max := 1, 0, 0
	for i := 0; i < len(digits); i++ {
		nextDigit := int(digits[i] - '0')
		if nextDigit == 0 {
			product, lenSb = 1, 0
			continue
		}
		if lenSb < span {
			product *= nextDigit
			lenSb++
		}
		if lenSb == span {
			if product > max {
				max = product
			}
			product /= int(digits[i-span+1] - '0')
			lenSb--
		}
	}
	return max, nil
}
