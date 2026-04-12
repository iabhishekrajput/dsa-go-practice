package stacks

// DailyTemperatures returns how many days until a warmer temperature
// for each day. If no future day is warmer, the value is 0.
//
// Example:
//
//	DailyTemperatures([]int{73,74,75,71,69,72,76,73}) → [1,1,4,2,1,1,0,0]
//
// Constraints:
//   - O(n) time, O(n) space
//
// Hint: Use a monotonic stack storing indices.
//
//	When current temp > stack top's temp, pop and record the difference.
func DailyTemperatures(temperatures []int) []int {
	// YOUR CODE HERE

	stack := []int{}
	result := make([]int, len(temperatures))

	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}

	return result
}
