package diffsquares

func sumOfFirstN(input int) int {
	return input * (input + 1) / 2
}

// SquareOfSum returns the square of the sun
func SquareOfSum(input int) int {
	sum := sumOfFirstN(input)
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first N natural numbers
func SumOfSquares(input int) int {
	return input * (input + 1) * (2*input + 1) / 6
}

// Difference returns the difference between the sum of the squares and the square of the sun
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
