// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"strings"
	"unicode"
)

// Remark identifies the type of remark this is a wrapper arround string
// it requires this structure to be able to use later the string methods
type Remark struct {
	remark string
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isShouting() bool {
	return strings.ToUpper(r.remark) == r.remark && strings.IndexFunc(r.remark, unicode.IsLetter) > -1
}

func remarkFactory(input string) Remark {
	return Remark{strings.TrimSpace(input)}
}

// Hey gets and answer from Bob
func Hey(remark string) string {
	r := remarkFactory(remark)
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isShouting() && r.isQuestion():
		return "Calm down, I know what I'm doing!"
	case r.isShouting():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
