package main

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	for _, v := range arr {
		println(v)
	}
	reverse(&arr)
	for _, v := range arr {
		println(v)
	}
}

// Rewrite reverse to use an array pointer instead of a slice.
func reverse(arr *[5]int) {
	l := len(*arr)

	for i := 0; i < l/2; i++ {
		// this already represents an item in the pointer
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}
