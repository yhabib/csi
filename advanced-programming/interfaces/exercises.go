package main

import (
	"fmt"
	"unsafe"
)

type integer struct {
	// pointer == 64b
	// tab  *itab
	tab  uintptr
	data unsafe.Pointer
}

func getInt(input interface{}) int {
	ptr := (*integer)(unsafe.Pointer(&input))
	return *(*int)(ptr.data)
}

func getMethods(input interface{}) {
}

func main() {
	var a interface{} = 2
	var b interface{} = 3

	fmt.Printf("are equal? %v \n", getInt(a) == 2)
	fmt.Printf("are equal? %v \n", getInt(b) == 3)
}
