package luhn

import (
	"unicode"
)

// Valid checks if given number satisfies Luhn formula
func Valid(n string) bool {
	var luhnSum int
	var digitCount int

	for i := len(n) - 1; i >= 0; i-- {
		r := rune(n[i])
		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		num := int(r - '0')
		digitCount++
		if digitCount%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		luhnSum += num
	}

	return digitCount > 1 && luhnSum%10 == 0
}
