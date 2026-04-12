package stacks

import "testing"

func TestMinStack(t *testing.T) {
	// Scenario 1: basic operations
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	if got := s.GetMin(); got != -3 {
		t.Errorf("GetMin() = %d, want -3", got)
	}
	s.Pop()
	if got := s.Top(); got != 0 {
		t.Errorf("Top() = %d, want 0", got)
	}
	if got := s.GetMin(); got != -2 {
		t.Errorf("GetMin() = %d, want -2", got)
	}

	// Scenario 2: all same values
	s2 := NewMinStack()
	s2.Push(5)
	s2.Push(5)
	s2.Push(5)
	if got := s2.GetMin(); got != 5 {
		t.Errorf("GetMin() = %d, want 5", got)
	}
	s2.Pop()
	if got := s2.GetMin(); got != 5 {
		t.Errorf("GetMin() after pop = %d, want 5", got)
	}

	// Scenario 3: descending then ascending
	s3 := NewMinStack()
	s3.Push(3)
	s3.Push(2)
	s3.Push(1)
	if got := s3.GetMin(); got != 1 {
		t.Errorf("GetMin() = %d, want 1", got)
	}
	s3.Pop()
	if got := s3.GetMin(); got != 2 {
		t.Errorf("GetMin() after pop = %d, want 2", got)
	}
	s3.Pop()
	if got := s3.GetMin(); got != 3 {
		t.Errorf("GetMin() after pop = %d, want 3", got)
	}
}
