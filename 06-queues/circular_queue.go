package queues

// CircularQueue implements a fixed-size circular queue (ring buffer).
type CircularQueue struct {
	// Define your fields here
	queue []int
	front int
	rear  int
	count int
	size int
}

// NewCircularQueue creates a circular queue with capacity k.
func NewCircularQueue(k int) CircularQueue {
	// Your solution here
	return CircularQueue{
		queue: make([]int, k),
		front: 0,
		rear: 0,
		count: 0,
		size: k,
	}
}

// Enqueue adds an element to the rear of the queue.
// Returns false if the queue is full.
func (q *CircularQueue) Enqueue(val int) bool {
	if q.count == q.size {
		return false
	}

	q.queue[q.rear] = val

	q.rear = (q.rear + 1) % q.size
	q.count++

	return true
}

// Dequeue removes an element from the front of the queue.
// Returns false if the queue is empty.
func (q *CircularQueue) Dequeue() bool {
	if q.count == 0 {
		return false
	}
	
	q.front = (q.front + 1) % q.size
	q.count--
	
	return true
}

// Front returns the front element. Returns -1 if empty.
func (q *CircularQueue) Front() int {
	if q.count == 0 {
		return -1
	}

	return q.queue[q.front]
}

// Rear returns the rear element. Returns -1 if empty.
func (q *CircularQueue) Rear() int {
	if q.count == 0 {
		return  -1
	}

	return q.queue[(q.rear - 1 + q.size) % q.size]
}

// IsEmpty returns true if the queue has no elements.
func (q *CircularQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull returns true if the queue is at capacity.
func (q *CircularQueue) IsFull() bool {
	return q.count == q.size
}
