package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.Push(3)
	stack.Push(5)
	description := "Add items to the Stack"
	expectedSize := 2
	if actualSize := stack.Size(); actualSize != expectedSize {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actualSize, expectedSize)
	}
	expectedPeek := 5
	if actualPeek := stack.Peek(); actualPeek != expectedPeek {
		t.Fatalf("FAIL: %s - s.Peek(): %d, expected %d", description, actualPeek, expectedPeek)
	}
	t.Logf("PASS: %s", description)
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push(5)
	description := "Remove item from top"
	expectedValue := 5
	if actual := stack.Pop(); actual != expectedValue {
		t.Fatalf("FAIL: %s - s.Pop(): %d, expected %d", description, actual, expectedValue)
	}
	if actual := stack.Size(); actual != 0 {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actual, expectedValue)
	}
	t.Logf("PASS: %s", description)
}

func TestSize(t *testing.T) {
	stack := Stack{}
	stack.Push(4)
	stack.Push("hello")
	description := "Size of stack"
	expected := 2
	if actual := stack.Size(); actual != expected {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestPeek(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	description := "Peek the top of the stack"
	expected := 1
	if actual := stack.Peek(); actual != expected {
		t.Fatalf("FAIL: %s - s.Push(1): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestIsEmpty(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	description := "Emptiness of Stack"
	expected := false
	if actual := stack.IsEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.IsEmpty(): %t, expected %t", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}
