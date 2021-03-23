package darts

// Notes:
// First iteration was with math.Hypo(). Quite useful but less performant
// By squaring the distances we avoid calculating the root of the point

// Score calculates the points for a given throw
func Score(x float64, y float64) int {
	d := x*x + y*y
	if d <= 1 {
		return 10
	} else if d <= 25 {
		return 5
	} else if d <= 100 {
		return 1
	}
	return 0
}
