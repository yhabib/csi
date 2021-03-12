// https://leetcode.com/problems/valid-parentheses/submissions/
package parentheses

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isClosing(r rune) bool {
	_, ok := pairs[r]
	return ok
}

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if isClosing(c) {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if opening := pairs[c]; opening != top {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
