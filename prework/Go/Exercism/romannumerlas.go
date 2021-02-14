package romannumerals

import (
	"fmt"
)

var (
	romans  = [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabics = [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

// ToRomanNumeral converts an arabic number into a roman one
func ToRomanNumeral(input int) (output string, err error) {
	if input < 1 || input > 3000 {
		err = fmt.Errorf("invalid number %d", input)
		return
	}
	for i := 0; i < len(arabics); i++ {
		for input >= arabics[i] {
			input -= arabics[i]
			output += romans[i]
		}
	}

	return
}
