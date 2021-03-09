// Package stack implements the Stack ADS
package stack

// Stack is a basic LIFO structure that resizes as needed
type Stack struct {
	items []interface{}
}

// Push appends to the top of the Stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns top of the stack
func (s *Stack) Pop() interface{} {
	size := len(s.items)
	if size == 0 {
		return nil
	}
	item := s.items[size-1]
	s.items = s.items[:size-1]
	return item
}

// Peek returns the element in the top of the Stack
func (s *Stack) Peek() interface{} {
	size := s.Size()
	if size == 0 {
		return nil
	}

	return s.items[size-1]
}

// IsEmpty returns if the Stack is empty or not
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns the size of the Stack
func (s *Stack) Size() int {
	return len(s.items)
}
