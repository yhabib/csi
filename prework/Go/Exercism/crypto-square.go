package cryptosquare

import (
	"math"
	"strings"
)

// Notes:
//  Iteration 1: Runes iteration        BenchmarkEncode-12        112322              9801 ns/op            7896 B/op        164 allocs/op
//  Iteration 2: String concatenation   BenchmarkEncode-12         60188             19821 ns/op            9888 B/op        663 allocs/op
//  Iteration 3: String builder         BenchmarkEncode-12        376096              3201 ns/op            1368 B/op         54 allocs/op
//

func getColAndRow(size int) (c int, r int) {
	c = int(math.Sqrt(float64(size)))
	r = c
	if r*c < size {
		c++
	}
	if r*c < size {
		r++
	}

	return
}

func isLowerLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func normalizer(r rune) rune {
	r = r | 0x20

	if !isLowerLetter(r) && !isDigit(r) {
		return -1
	}
	return r
}

// Encode transforms the given input applying
func Encode(input string) string {
	normalizedInput := strings.Map(normalizer, input)
	l := len(normalizedInput)
	c, r := getColAndRow(l)

	var b strings.Builder

	for i := 0; i < c; i++ {
		if i > 0 {
			b.WriteByte(' ')
		}
		for j := 0; j < r*c; j += c {
			if i+j < l {
				b.WriteByte(normalizedInput[i+j])
			} else {
				b.WriteByte(' ')
			}
		}
	}

	return b.String()
}
