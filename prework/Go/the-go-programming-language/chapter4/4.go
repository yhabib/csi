package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s, 4)
	fmt.Println(s)
}

// Write a version of rotate that operates in a single pass.
func rotate(arr []int, amount int) {
	reverse(arr[:amount])
	reverse(arr[amount:])
	reverse(arr)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
