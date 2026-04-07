package queues

import "testing"

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()

	// Basic FIFO
	q.Push(1)
	q.Push(2)

	if got := q.Peek(); got != 1 {
		t.Errorf("Peek() = %d, want 1", got)
	}
	if got := q.Pop(); got != 1 {
		t.Errorf("Pop() = %d, want 1", got)
	}
	if got := q.Empty(); got != false {
		t.Errorf("Empty() = %v, want false", got)
	}

	// Interleaved push and pop
	q2 := NewMyQueue()
	q2.Push(10)
	q2.Push(20)
	if got := q2.Pop(); got != 10 {
		t.Errorf("Pop() = %d, want 10", got)
	}
	q2.Push(30)
	if got := q2.Pop(); got != 20 {
		t.Errorf("Pop() = %d, want 20", got)
	}
	if got := q2.Pop(); got != 30 {
		t.Errorf("Pop() = %d, want 30", got)
	}
	if got := q2.Empty(); got != true {
		t.Errorf("Empty() = %v, want true", got)
	}

	// Many elements
	q3 := NewMyQueue()
	for i := 1; i <= 5; i++ {
		q3.Push(i)
	}
	for i := 1; i <= 5; i++ {
		if got := q3.Pop(); got != i {
			t.Errorf("Pop() = %d, want %d", got, i)
		}
	}
}
