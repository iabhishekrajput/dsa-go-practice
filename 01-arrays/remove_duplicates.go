package arrays

// RemoveDuplicates removes duplicates from a sorted slice in-place.
// Returns the number of unique elements.
//
// Example:
//   nums := []int{1, 1, 2, 3, 3, 3, 4}
//   k := RemoveDuplicates(nums) // k = 4, nums[:k] = [1, 2, 3, 4]
//
// Constraints:
//   - The slice is sorted in non-decreasing order
//   - Modify the slice in-place (no extra slice allowed)
//   - O(1) extra space
//
// Hint: Use two pointers — a "slow" writer and a "fast" reader.
func RemoveDuplicates(nums []int) int {
	// YOUR CODE HERE
	
	if len(nums) == 0 {
		return 0
	}

	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}

	return left + 1
}
