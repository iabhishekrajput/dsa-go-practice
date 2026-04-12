package queues

// MaxSlidingWindow returns the max value in each sliding window of size k.
//
// Example:
//   MaxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3) → [3,3,5,5,6,7]
//
// Constraints:
//   - O(n) time
//
// Hint: Use a deque ([]int) storing indices in decreasing value order.
//       Front = current max. Pop back if new value is bigger.
//       Pop front if index is outside the window.
func MaxSlidingWindow(nums []int, k int) []int {
	// YOUR CODE HERE

	deque := make([]int, 0, len(nums))
	result := []int{}

	for idx, num := range nums {
		for len(deque) > 0 && deque[0] <= idx - k {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque) - 1]] <= num {
			deque = deque[:len(deque) - 1]
		}

		deque = append(deque, idx)

		if idx >= k - 1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
