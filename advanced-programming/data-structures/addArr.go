package main

import (
	"fmt"
	"unsafe"
)

// in this machine int is 64b
// to go through the array +8 ?? 64b / 8b = 8 bytes -> every 8 bytes I have a another number
// what do I have every byte then? random data? old data?
func addArr(arr []int) (sum int) {
	ptr := *(*uintptr)(unsafe.Pointer(&arr))
	maxPtr := ptr + uintptr(8*len(arr))
	for i := ptr; i < maxPtr; i += 8 {
		sum += *(*int)(unsafe.Pointer(i))
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Printf("%d == %d\n", sum, addArr(arr))
}
