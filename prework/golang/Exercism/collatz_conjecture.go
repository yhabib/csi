package collatzconjecture

import "errors"

// CollatzConjecture returns
func CollatzConjecture(n int) (out int, ok error) {
	if n < 1 {
		return out, errors.New("n is less than 1")
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			out++
		} else {
			n = 3*n + 1
			out++
		}
	}
	return out, nil
}
