package queues

// MyQueue implements a FIFO queue using two stacks.
//
// Hint: Use inStack for pushes, outStack for pops.
//       When outStack is empty, pour all of inStack into it.
type MyQueue struct {
	// YOUR CODE HERE
	inQueue []int
	outQueue []int
}

func NewMyQueue() MyQueue {
	// YOUR CODE HERE
	return MyQueue{
		inQueue: []int{},
		outQueue: []int{},
	}
}

func (q *MyQueue) Push(val int) {
	// YOUR CODE HERE
	q.inQueue = append(q.inQueue, val)
}

func (q *MyQueue) Pop() int {
	// YOUR CODE HERE
	top := q.Peek()
	q.outQueue = q.outQueue[:len(q.outQueue) - 1]

	return top
}

func (q *MyQueue) Peek() int {
	// YOUR CODE HERE
	if len(q.outQueue) == 0 {
		for i := len(q.inQueue) - 1; i >= 0; i-- {
			q.outQueue = append(q.outQueue, q.inQueue[i])
		}
		q.inQueue = q.inQueue[:0]
	}

	return q.outQueue[len(q.outQueue) - 1]
}

func (q *MyQueue) Empty() bool {
	// YOUR CODE HERE
	return len(q.inQueue) == 0 && len(q.outQueue) == 0
}
