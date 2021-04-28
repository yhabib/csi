package main

import (
	"fmt"
	"unsafe"
)

// Maps bit representation from float64 into a uint64
func fromFloatToUint(num float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&num))
}

func main() {
	var f float64 = 1.0
	ui := fromFloatToUint(f)

	fmt.Printf("float64 bit pattern: %064b\n", &f)
	fmt.Printf("uint64 bit pattern: %064b\n", ui)
	fmt.Printf("uint64 value: %d\n", ui)
}

// Output, not sure what I should see
// 		float64 bit pattern: 0000000000000000000000001100000000000000000000010100000011001000
// 		uint64 bit pattern: 0011111111110000000000000000000000000000000000000000000000000000
// 		uint64 value: 4607182418800017408
