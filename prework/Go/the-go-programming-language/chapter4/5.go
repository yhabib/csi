package main

import "fmt"

func main() {
	slice := []string{"hello", "woordl", "woordl"}
	fmt.Println(slice)
	fmt.Println(removeAdjacents(slice))
}

// Write an in-place function to eliminate adjacent duplicates in a []string slice.
func removeAdjacents(slice []string) []string {
	out := slice[:0] // Zero size slice of the original, by doing this we save memory as we reuse the original one
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != slice[i+1] {
			out = append(out, slice[i])
		}
	}
	return out
}
