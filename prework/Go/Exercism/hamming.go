import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var distance int
	// Expensive operation
	runeA := []rune(a)
	runeB := []rune(b)

	if len(runeA) != len(runeB) {
		return distance, errors.New("inputs have different length")
	}

	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			distance++
		}
	}
	return distance, nil
}