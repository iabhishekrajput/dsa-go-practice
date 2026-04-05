package arrays

// TwoSum returns indices of two numbers in nums that add up to target.
// Exactly one solution exists. You may not use the same element twice.
//
// Example:
//   TwoSum([]int{2, 7, 11, 15}, 9) → [0, 1]
//
// Challenge: Solve it in O(n) time.
//
// Hint: Use a map to remember what you've seen.
//       For each number, the "complement" you need is (target - number).
//       Have you seen that complement before?
func TwoSum(nums []int, target int) []int {
	// YOUR CODE HERE

	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		comp := target - v
		idx, ok := seen[comp]; if ok {
			return []int{i, idx}
		}

		seen[v] = i
	}

	return nil
}
