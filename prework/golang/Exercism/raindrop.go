package raindrops

import "fmt"

// Convert returns the passed argument as string
func Convert(num int) string {
	var message string

	if num%3 == 0 {
		message += "Pling"
	}
	if num%5 == 0 {
		message += "Plang"
	}
	if num%7 == 0 {
		message += "Plong"
	}

	if len(message) == 0 {
		return fmt.Sprint(num)
	}

	return message
}
