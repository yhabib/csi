package accumulate

// In place solution

// Accumulate returns the the input slice after converter operation has been executed
func Accumulate(input []string, converter func(string) string) []string {
	output := input[:0]
	for _, str := range input {
		output = append(output, converter(str))
	}
	return output
}

// Alternative 1
// func Accumulate(input []string, converter func(string) string) []string {
// 	for i, str := range input {
// 		input[i] = converter(str)
// 	}
// 	return input
// }

// Alternative 2
// func Accumulate(input []string, converter func(string) string) (output []string) {
// 	for _, str := range input {
// 		output = append(output, converter(str))
// 	}
// 	return
// }
