package stacks

// MinStack is a stack that supports Push, Pop, Top, and GetMin in O(1).
//
// Hint: Use two slices — one for values, one tracking the min at each level.
type MinStack struct {
	// YOUR CODE HERE
	stack []int
	minStack []int
}

func NewMinStack() MinStack {
	// YOUR CODE HERE
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	// YOUR CODE HERE
	if len(s.minStack) == 0 {
		s.minStack = append(s.minStack, val)
	} else {
		s.minStack = append(s.minStack, min(val, s.GetMin()))
	}
	s.stack = append(s.stack, val)
}

func (s *MinStack) Pop() {
	// YOUR CODE HERE
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack) - 1]
	}

	if len(s.minStack) > 0 {
		s.minStack = s.minStack[:len(s.minStack) - 1]
	}
}

func (s *MinStack) Top() int {
	// YOUR CODE HERE
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack) - 1]
}

func (s *MinStack) GetMin() int {
	// YOUR CODE HERE
	if len(s.minStack) == 0 {
		return 0
	}

	return s.minStack[len(s.minStack) - 1]
}
