package main

import (
	"fmt"
	"unsafe"
)

var 

// An unsafe.Pointer may also be converted to a uintptr that holds
// the pointer’s numeric value, letting us perform arithmetic on addresses.
// but then what is &str?? a pointer to the string struct of two words?
func isSameLocation(a string, b string) bool {
	// &a is the address of the variable in the Stack
	// fmt.Println(&a)
	// up holds the value of &a that is the address of the variable in the Stack up -> at the end is an int
	// &up is the address of this new variable in the stack that holds the value of &a
	// up := unsafe.Pointer(&a)
	// fmt.Println(up)
	// fmt.Println(&up)
	// Casts the unsafe pointer to a pointer of uintptr, same value as unsafe.Pointer(&a) both ints
	fmt.Println((*uintptr)(unsafe.Pointer(&a)))
	// Casts the unsafe pointer to a uintptr
	fmt.Println(uintptr(unsafe.Pointer(&a)))
	fmt.Println(*(*uintptr)(unsafe.Pointer(&a)))
	return *(*uintptr)(unsafe.Pointer(&a)) == *(*uintptr)(unsafe.Pointer(&b))
}

func main() {
	a := "hello"
	b := "hello"
	fmt.Printf("is same location? %t \n", isSameLocation(a, b))
}
