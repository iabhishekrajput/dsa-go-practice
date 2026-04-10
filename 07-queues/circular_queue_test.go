package queues

import "testing"

func TestCircularQueue_BasicOps(t *testing.T) {
	q := NewCircularQueue(3)

	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	if q.IsFull() {
		t.Error("new queue should not be full")
	}
	if q.Front() != -1 {
		t.Errorf("Front() on empty queue = %d, want -1", q.Front())
	}
	if q.Rear() != -1 {
		t.Errorf("Rear() on empty queue = %d, want -1", q.Rear())
	}
}

func TestCircularQueue_EnqueueDequeue(t *testing.T) {
	q := NewCircularQueue(3)

	if !q.Enqueue(1) {
		t.Error("Enqueue(1) should succeed")
	}
	if !q.Enqueue(2) {
		t.Error("Enqueue(2) should succeed")
	}
	if !q.Enqueue(3) {
		t.Error("Enqueue(3) should succeed")
	}
	if q.Enqueue(4) {
		t.Error("Enqueue(4) should fail — queue is full")
	}

	if q.Front() != 1 {
		t.Errorf("Front() = %d, want 1", q.Front())
	}
	if q.Rear() != 3 {
		t.Errorf("Rear() = %d, want 3", q.Rear())
	}
	if !q.IsFull() {
		t.Error("queue should be full")
	}
}

func TestCircularQueue_WrapAround(t *testing.T) {
	q := NewCircularQueue(3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue 1, making room at the "beginning"
	if !q.Dequeue() {
		t.Error("Dequeue should succeed")
	}
	if q.Front() != 2 {
		t.Errorf("Front() after dequeue = %d, want 2", q.Front())
	}

	// Enqueue 4 — should wrap around to index 0
	if !q.Enqueue(4) {
		t.Error("Enqueue(4) should succeed after dequeue")
	}
	if q.Rear() != 4 {
		t.Errorf("Rear() = %d, want 4", q.Rear())
	}

	// Verify order: 2, 3, 4
	if q.Front() != 2 {
		t.Errorf("Front() = %d, want 2", q.Front())
	}
}

func TestCircularQueue_DequeueEmpty(t *testing.T) {
	q := NewCircularQueue(2)

	if q.Dequeue() {
		t.Error("Dequeue on empty queue should return false")
	}

	q.Enqueue(10)
	q.Dequeue()

	if q.Dequeue() {
		t.Error("Dequeue on emptied queue should return false")
	}
	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing all elements")
	}
}

func TestCircularQueue_FullCycle(t *testing.T) {
	q := NewCircularQueue(2)

	// Fill and drain multiple times to stress wrap-around
	for cycle := 0; cycle < 5; cycle++ {
		if !q.Enqueue(cycle*10 + 1) {
			t.Errorf("cycle %d: Enqueue 1 failed", cycle)
		}
		if !q.Enqueue(cycle*10 + 2) {
			t.Errorf("cycle %d: Enqueue 2 failed", cycle)
		}
		if q.Enqueue(cycle*10 + 3) {
			t.Errorf("cycle %d: Enqueue 3 should fail (full)", cycle)
		}

		if q.Front() != cycle*10+1 {
			t.Errorf("cycle %d: Front() = %d, want %d", cycle, q.Front(), cycle*10+1)
		}
		if q.Rear() != cycle*10+2 {
			t.Errorf("cycle %d: Rear() = %d, want %d", cycle, q.Rear(), cycle*10+2)
		}

		q.Dequeue()
		q.Dequeue()
	}

	if !q.IsEmpty() {
		t.Error("should be empty after full cycles")
	}
}

func TestCircularQueue_SingleCapacity(t *testing.T) {
	q := NewCircularQueue(1)

	if !q.Enqueue(42) {
		t.Error("Enqueue should succeed")
	}
	if !q.IsFull() {
		t.Error("should be full")
	}
	if q.Front() != 42 {
		t.Errorf("Front() = %d, want 42", q.Front())
	}
	if q.Rear() != 42 {
		t.Errorf("Rear() = %d, want 42", q.Rear())
	}
	if !q.Dequeue() {
		t.Error("Dequeue should succeed")
	}
	if !q.IsEmpty() {
		t.Error("should be empty")
	}
}
