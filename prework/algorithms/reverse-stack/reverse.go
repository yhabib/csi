package reverse

import (
	"strings"

	"github.com/yhabib/go-modules/stack"
)

// Reverse returns the reversed version of given string
func Reverse(s string) string {
	var reversed strings.Builder
	stack := stack.Stack{}
	for _, c := range s {
		stack.Push(c)
	}
	for !stack.IsEmpty() {
		// Type assertion
		reversed.WriteRune(stack.Pop().(rune))
	}
	return reversed.String()
}
