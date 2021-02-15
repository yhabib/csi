// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (output string) {
	isNew := true
	for _, w := range s {
		if isNew && unicode.IsLetter(w) {
			output += strings.ToUpper(string(w))
			isNew = false
		}
		if unicode.IsSpace(w) || w == '_' || w == '-' {
			isNew = true
		}
	}
	return
}

// Better approach
// func Abbreviate(s string) (out string) {
// 	s = strings.Replace(s, "-", " ", -1)
// 	words := strings.Fields(s)
// 	for i := range words {
// 		out += string(words[i][0])
// 	}
// 	return strings.ToUpper(out)
// }
