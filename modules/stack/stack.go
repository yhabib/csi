package stack

// ADS
//  * New: 		 Constructor
//  * Push: 	 Append item to the top of the stack. Return nothing
//  * Pop: 		 Remove item from the top of the stack. Return item
//  * Peek: 	 Check item on the top of the stack. Return item.
//  * IsEmpty: Checks if the Stack contains items
//  * Size:    Get the size of the stack. Return size

// Stack is a basic LIFO structure that resizes as needed
type Stack struct {
	Stack []interface{}
}

// New is the constructor
func (s *Stack) New() {
	s.Stack = make([]interface{}, 0)
}

// Push appends to the top of the Stack
func (s *Stack) Push(item interface{}) {
	// a slice referencing the storage of s.Stack
	s.Stack = append(s.Stack[:], item)
}

// Pop removes and returns top of the stack
func (s *Stack) Pop() interface{} {
	size := len(s.Stack)
	if size == 0 {
		return nil
	}
	item := s.Stack[size-1]
	s.Stack = s.Stack[:size-1]
	return item
}

// Peek returns the element in the top of the Stack
func (s *Stack) Peek() interface{} {
	size := len(s.Stack)
	if size == 0 {
		return nil
	}

	return s.Stack[size-1]
}

// IsEmpty returns if the Stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}

// Size returns the size of the Stack
func (s *Stack) Size() int {
	return len(s.Stack)
}
