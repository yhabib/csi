package strain

type (
	// Ints is a collection of ints
	Ints []int
	// Lists is a collection of collections of ints
	Lists [][]int
	// Strings is a collection of strings
	Strings []string
)

// Keep keeps elements based on predicate
func (ints Ints) Keep(predicate func(int) bool) (newInts Ints) {
	for _, v := range ints {
		if predicate(v) {
			newInts = append(newInts, v)
		}
	}
	return
}

// Discard discards elements based on predicate
func (ints Ints) Discard(predicate func(int) bool) (new Ints) {
	return ints.Keep(func(v int) bool { return !predicate(v) })
}

// Keep keeps elements based on predicate
func (lists Lists) Keep(predicate func([]int) bool) (newLists Lists) {
	for _, v := range lists {
		if predicate(v) {
			newLists = append(newLists, v)
		}
	}
	return
}

// Keep keeps elements based on predicate
func (strings Strings) Keep(predicate func(string) bool) (newStrings Strings) {
	for _, v := range strings {
		if predicate(v) {
			newStrings = append(newStrings, v)
		}
	}
	return
}
