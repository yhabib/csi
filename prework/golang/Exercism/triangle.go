// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind string

// const (
// 	NaT = "NaT" // not a triangle
// 	Equ = "Equ" // equilateral
// 	Iso = "Iso" // isosceles
// 	Sca = "Sca" // scalene
// )
type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	product := a * b * c
	if product == 0 || math.IsInf(product, 0) || product != product {
		k = NaT
	} else if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
