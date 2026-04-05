package arrays

// MaxSumSubarray returns the maximum sum among all contiguous subarrays of size k.
//
// Example:
//   MaxSumSubarray([]int{2, 1, 5, 1, 3, 2}, 3) → 9
//
// Constraints:
//   - 1 <= k <= len(nums)
//   - O(n) time, O(1) space
//
// Hint: Build the sum of the first window (0..k-1).
//       Then slide: add nums[i], subtract nums[i-k].
//       Track the max as you go.
func MaxSumSubarray(nums []int, k int) int {
	// YOUR CODE HERE

	sums := 0
	for i := range k {
		sums += nums[i]
	}

	maxSum := sums

	for i, j := 0, k; j < len(nums); i, j = i + 1, j + 1 {
		sums = sums - nums[i] + nums[j]
		if sums > maxSum {
			maxSum = sums
		}
	} 

	return maxSum
}
