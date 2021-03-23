package reverse

// Notes:
// First approach  for _, c := range in { out = string(c) + out	}	return
// This is less performant: more iterations and one cast per iteration

// Reverse returns the reversed version of the provided string in place
func Reverse(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
