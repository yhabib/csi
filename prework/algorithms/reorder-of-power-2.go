import "strconv"

func countOfDigits(n int) []int {
	countOfDigits := make([]int, 10)
	str := strconv.Itoa(n)
	for _, d := range str {
		countOfDigits[int(d-'0')]++
	}
	return countOfDigits
}

func reorderedPowerOf2(N int) bool {
	var mask int
	countOfN := countOfDigits(N)

ex:
	for i := 0; i < 32; i++ {
		mask = 1 << i
		countOfPowOf2 := countOfDigits(mask)
		for j := 0; j < len(countOfN); j++ {
			if countOfN[j] != countOfPowOf2[j] {
				continue ex
			}
		}
		return true
	}
	return false
}