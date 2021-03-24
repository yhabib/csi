package grains

import (
	"errors"
)

// Notes:
//  Initially used ~bs for the Total because I didn't realize that I could achieve it with 1<<64 - 1 😅
// 	No need the extra step of | against zero, I can just shift the bit the required amount

// Total returns the total number of grains
func Total() (bs uint64) {
	return 1<<64 - 1
}

// Square calculates the amount of grain for given input
func Square(input int) (bs uint64, err error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input is out of range")
	}

	return 1 << (input - 1), nil
}
