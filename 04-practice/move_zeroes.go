package practice

// MoveZeroes moves all 0s to the end of the slice in-place,
// maintaining the relative order of non-zero elements.
//
// Example:
//
//	nums := []int{0, 1, 0, 3, 12}
//	MoveZeroes(nums) // nums = [1, 3, 12, 0, 0]
//
// Constraints:
//   - In-place, no extra slice
//   - O(n) time
func MoveZeroes(nums []int) {
	// YOUR CODE HERE

	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
