package stack

import (
	"testing"
)

var newStack = Stack{}

func TestSize(t *testing.T) {
	description := "Size of empty stack"
	expected := 0
	if actual := newStack.Size(); actual != expected {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestPeek(t *testing.T) {
	newStack.Push(1)
	description := "Peek the top of the stack"
	expected := 1
	if actual := newStack.Peek(); actual != expected {
		t.Fatalf("FAIL: %s - s.Push(1): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestPush(t *testing.T) {
	newStack.Push(3)
	description := "Add item to Stack"
	expectedSize := 2
	if actualSize := newStack.Size(); actualSize != expectedSize {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actualSize, expectedSize)
	}
	expectedPeek := 3
	if actualPeek := newStack.Peek(); actualPeek != expectedPeek {
		t.Fatalf("FAIL: %s - s.Peek(): %d, expected %d", description, actualPeek, expectedPeek)
	}
	t.Logf("PASS: %s", description)
}

func TestPop(t *testing.T) {
	newStack.Push(5)
	description := "Remove item from top"
	expectedValue := 5
	if actual := newStack.Pop(); actual != expectedValue {
		t.Fatalf("FAIL: %s - s.Pop(): %d, expected %d", description, actual, expectedValue)
	}
	if actual := newStack.Size(); actual != 2 {
		t.Fatalf("FAIL: %s - s.Size(): %d, expected %d", description, actual, expectedValue)
	}
	t.Logf("PASS: %s", description)
}

func TestIsEmpty(t *testing.T) {
	description := "Emptiness of Stack"
	expected := false
	if actual := newStack.IsEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.IsEmpty(): %t, expected %t", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}
