// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb returns a string containing a proverb
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	var proverb []string
	if length == 0 {
		return proverb
	}

	for i := 0; length > 1 && i < length-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	return append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
