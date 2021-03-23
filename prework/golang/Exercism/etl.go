package etl

import (
	"strings"
)

// Transform performs ETP operation
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for points, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = points
		}
	}

	return output
}
