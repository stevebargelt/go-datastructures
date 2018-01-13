package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Errorf("Length of an New() empty stack should be 0")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if s.Peek().(int) != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Pop().(int) != 1 {
		t.Errorf("Top item should have been 1")
	}

	if s.Len() != 0 {
		t.Errorf("Popped the only item, Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if s.Peek().(int) != 2 {
		t.Errorf("Top of the stack should be 2")
	}
	s.Pop()
	s.Pop()

	s.Push(0)
	s.Push(1)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(5)
	expected := 12
	// Check iteration
	sum := 0
	initialLen := s.Len()
	for i := 1; i < initialLen; i++ {
		sum += s.Pop().(int)
		fmt.Printf("Sum = %v\n", sum)
	}
	if sum != expected {
		t.Errorf("sum over s = %d, expected 12", sum)
	}

}
