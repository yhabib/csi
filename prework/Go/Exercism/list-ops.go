package listops

// Why functions are exported even those lower case ??

type (
	binFunc   func(int, int) int
	predFunc  func(int) bool
	unaryFunc func(int) int

	// IntList represents a list of integers
	IntList []int
)

// Foldl reduces list from the left
func (l IntList) Foldl(fn binFunc, acc int) int {
	for _, item := range l {
		acc = fn(acc, item)
	}
	return acc
}

// Foldr reduces list from the right
func (l IntList) Foldr(fn binFunc, acc int) int {
	for _, item := range l.Reverse() {
		acc = fn(item, acc)
	}
	return acc
}

// Filter filters list
func (l IntList) Filter(fn predFunc) IntList {
	nl := make(IntList, 0, l.Length())
	for _, item := range l {
		if fn(item) {
			nl = append(nl, item)
		}
	}
	return nl
}

// Length returns size of list
func (l IntList) Length() (size int) {
	for ; size < len(l); size++ {
	}
	return size
}

// Map applies op in place to all items in list
func (l IntList) Map(fn unaryFunc) IntList {
	for i, item := range l {
		l[i] = fn(item)
	}
	return l
}

// Reverse reverses list in place
func (l IntList) Reverse() IntList {
	size := l.Length()
	for i := 0; i < size/2; i++ {
		l[i], l[size-1-i] = l[size-1-i], l[i]
	}
	return l
}

// Append list into end of first list
func (l IntList) Append(list IntList) IntList {
	f := make(IntList, 0, len(l)+len(list))
	f = append(f, l...)
	f = append(f, list...)
	return f
}

// Concat two slices
func (l IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		l = l.Append(list)
	}
	return l
}
